// Copyright (c) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package fs

import (
	"fmt"
	"os"
	"path/filepath"

	persistapi "github.com/kata-containers/kata-containers/src/runtime/virtcontainers/persist/api"
)

var mockTesting = false

type MockFS struct {
	// inherit from FS. Overwrite if needed.
	*FS
}

func EnableMockTesting() {
	mockTesting = true
}

func MockStorageRootPath() string {
	return filepath.Join(os.TempDir(), "vc", "mockfs")
}

func MockRunStoragePath() string {
	return filepath.Join(MockStorageRootPath(), sandboxPathSuffix)
}

func MockRunVMStoragePath() string {
	return filepath.Join(MockStorageRootPath(), vmPathSuffix)
}

func MockFSInit(rootPath string) (persistapi.PersistDriver, error) {
	driver, err := Init()
	if err != nil {
		return nil, fmt.Errorf("Could not create Mock FS driver: %v", err)
	}

	fsDriver, ok := driver.(*FS)
	if !ok {
		return nil, fmt.Errorf("Could not create Mock FS driver")
	}

	fsDriver.storageRootPath = rootPath
	fsDriver.driverName = "mockfs"

	return &MockFS{fsDriver}, nil
}

func MockAutoInit() (persistapi.PersistDriver, error) {
	if mockTesting {
		return MockFSInit(MockStorageRootPath())
	}
	return nil, nil
}
