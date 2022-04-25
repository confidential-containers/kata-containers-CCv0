#!/bin/bash
#
# Copyright (c) 2022 IBM
#
# SPDX-License-Identifier: Apache-2.0

set -o errexit
set -o nounset
set -o pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

source "${script_dir}/../../scripts/lib.sh"

# disabling set -u because scripts attempt to expand undefined variables
set +u
ovmf_repo="${ovmf_repo:-}"
ovmf_dir="edk2"
ovmf_version="${ovmf_version:-}"
kata_version="${kata_version:-}"
DESTDIR="${DESTDIR:-../../destdir}"
PREFIX="${PREFIX:-/opt/kata}"
ovmf_build="${ovmf_build:-sev}"
architecture="X64"
toolchain="GCC5"
build_target="RELEASE"

if [ -z "$ovmf_repo" ]; then
       info "Get ovmf information from runtime versions.yaml"
       ovmf_repo=$(get_from_kata_deps "externals.ovmf.url" "${kata_version}")
fi
[ -n "$ovmf_repo" ] || die "failed to get ovmf repo"

if [ "${ovmf_build}" == "sev" ]; then
       [ -n "$ovmf_version" ] || ovmf_version=$(get_from_kata_deps "externals.ovmf.sev.version" "${kata_version}")
       [ -n "$ovmf_package" ] || ovmf_package=$(get_from_kata_deps "externals.ovmf.sev.package" "${kata_version}")
       [ -n "$package_output_dir" ] || package_output_dir=$(get_from_kata_deps "externals.ovmf.sev.package_output_dir" "${kata_version}")
elif [ "${ovmf_build}" == "x64" ]; then
       [ -n "$ovmf_version" ] || ovmf_version=$(get_from_kata_deps "externals.ovmf.x64.version" "${kata_version}")
       [ -n "$ovmf_package" ] || ovmf_package=$(get_from_kata_deps "externals.ovmf.x64.package" "${kata_version}")
       [ -n "$package_output_dir" ] || package_output_dir=$(get_from_kata_deps "externals.ovmf.x64.package_output_dir" "${kata_version}")

fi
[ -n "$ovmf_version" ] || die "failed to get ovmf version or commit"
[ -n "$ovmf_package" ] || die "failed to get ovmf package or commit"
[ -n "$package_output_dir" ] || die "failed to get ovmf package or commit"

info "Build ${ovmf_repo} version: ${ovmf_version}"

build_root=$(mktemp -d)
pushd $build_root
git clone "${ovmf_repo}"
cd "${ovmf_dir}"
git checkout "${ovmf_version}"
git submodule init
git submodule update

info "Using BaseTools make target"
make -C BaseTools/

info "Calling edksetup script"
source edksetup.sh

if [ "${ovmf_build}" == "sev" ]; then
       info "Creating dummy grub file"
       #required for building AmdSev package without grub
       touch OvmfPkg/AmdSev/Grub/grub.efi
fi

info "Building ovmf"
build -b "${build_target}" -t "${toolchain}" -a "${architecture}" -p "${ovmf_package}"

info "Done Building"
pwd

build_path="Build/${package_output_dir}/${build_target}_${toolchain}/FV/OVMF.fd"
stat "${build_path}"

#need to leave tmp dir
popd

info "Install fd to destdir"
mkdir -p "$DESTDIR/$PREFIX/share/ovmf"
cp $build_root/$ovmf_dir/"${build_path}" "$DESTDIR/$PREFIX/share/ovmf"

cleanup() {
    if [ "${DEBUG:-}" == true ]; then
       info "Not deleted build root directory: ${build_root}"
       return
    fi
    rm -rf "${build_root}"
}

trap cleanup EXIT