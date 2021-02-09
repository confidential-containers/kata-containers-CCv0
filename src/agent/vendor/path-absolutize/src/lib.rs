extern crate path_dedot;

#[cfg(windows)]
#[macro_use]
extern crate slash_formatter;

use std::io;
use std::path::{Path, PathBuf};

#[doc(hidden)]
pub use path_dedot::CWD;

pub use path_dedot::MAIN_SEPARATOR;

pub use path_dedot::update_cwd;

/// Make `Path` and `PathBuf` have `absolutize` and `absolutize_virtually` method.
pub trait Absolutize {
    /// Get an absolute path. This works even if the path does not exist.
    ///
    /// Please read the following examples to know the parsing rules.
    ///
    /// # Examples
    ///
    /// The dots in a path will be parsed even if it is already an absolute path (which means the path starts with a `MAIN_SEPARATOR` on Unix-like systems).
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::path::Path;
    ///
    /// use path_absolutize::*;
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("/path/to/123/456");
    ///
    ///     assert_eq!("/path/to/123/456", p.absolutize().unwrap().to_str().unwrap());
    /// }
    /// ```
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::path::Path;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("/path/to/./123/../456");
    ///
    ///     assert_eq!("/path/to/456", p.absolutize().unwrap().to_str().unwrap());
    /// }
    /// ```
    ///
    /// If a path starts with a single dot, the dot means **current working directory**.
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::env;
    /// use std::path::Path;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("./path/to/123/456");
    ///
    ///     assert_eq!(
    ///         Path::join(env::current_dir().unwrap().as_path(), Path::new("path/to/123/456"))
    ///             .to_str()
    ///             .unwrap(),
    ///         p.absolutize().unwrap().to_str().unwrap()
    ///     );
    /// }
    /// ```
    ///
    /// If a path starts with a pair of dots, the dots means the parent of **current working directory**. If **current working directory** is **root**, the parent is still **root**.
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::env;
    /// use std::path::Path;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("../path/to/123/456");
    ///
    ///     let cwd = env::current_dir().unwrap();
    ///
    ///     let cwd_parent = cwd.parent();
    ///
    ///     match cwd_parent {
    ///         Some(cwd_parent) => {
    ///             assert_eq!(
    ///                 Path::join(&cwd_parent, Path::new("path/to/123/456")).to_str().unwrap(),
    ///                 p.absolutize().unwrap().to_str().unwrap()
    ///             );
    ///         }
    ///         None => {
    ///             assert_eq!(
    ///                 Path::join(Path::new("/"), Path::new("path/to/123/456")).to_str().unwrap(),
    ///                 p.absolutize().unwrap().to_str().unwrap()
    ///             );
    ///         }
    ///     }
    /// }
    /// ```
    ///
    /// A path which does not start with a `MAIN_SEPARATOR`, **Single Dot** and **Double Dots**, will act like having a single dot at the start when `absolutize` method is used.
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::env;
    /// use std::path::Path;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("path/to/123/456");
    ///
    ///     assert_eq!(
    ///         Path::join(env::current_dir().unwrap().as_path(), Path::new("path/to/123/456"))
    ///             .to_str()
    ///             .unwrap(),
    ///         p.absolutize().unwrap().to_str().unwrap()
    ///     );
    /// }
    /// ```
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::env;
    /// use std::path::Path;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("path/../../to/123/456");
    ///
    ///     let cwd = env::current_dir().unwrap();
    ///
    ///     let cwd_parent = cwd.parent();
    ///
    ///     match cwd_parent {
    ///         Some(cwd_parent) => {
    ///             assert_eq!(
    ///                 Path::join(&cwd_parent, Path::new("to/123/456")).to_str().unwrap(),
    ///                 p.absolutize().unwrap().to_str().unwrap()
    ///             );
    ///         }
    ///         None => {
    ///             assert_eq!(
    ///                 Path::join(Path::new("/"), Path::new("to/123/456")).to_str().unwrap(),
    ///                 p.absolutize().unwrap().to_str().unwrap()
    ///             );
    ///         }
    ///     }
    /// }
    /// ```
    fn absolutize(&self) -> io::Result<PathBuf>;

    /// Get an absolute path **only under a specific directory**. This works even if the path does not exist.
    ///
    /// Please read the following examples to know the parsing rules.
    ///
    /// # Examples
    ///
    /// The dots in a path will be parsed even if it is already an absolute path (which means the path starts with a `MAIN_SEPARATOR` on Unix-like systems).
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::path::Path;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("/path/to/123/456");
    ///
    ///     assert_eq!("/path/to/123/456", p.absolutize_virtually("/").unwrap().to_str().unwrap());
    /// }
    /// ```
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::path::Path;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("/path/to/./123/../456");
    ///
    ///     assert_eq!("/path/to/456", p.absolutize_virtually("/").unwrap().to_str().unwrap());
    /// }
    /// ```
    ///
    /// Every absolute path should under the virtual root.
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::path::Path;
    ///
    /// use std::io::ErrorKind;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("/path/to/123/456");
    ///
    ///     assert_eq!(
    ///         ErrorKind::InvalidInput,
    ///         p.absolutize_virtually("/virtual/root").unwrap_err().kind()
    ///     );
    /// }
    /// ```
    ///
    /// Every relative path should under the virtual root.
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::path::Path;
    ///
    /// use std::io::ErrorKind;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("./path/to/123/456");
    ///
    ///     assert_eq!(
    ///         ErrorKind::InvalidInput,
    ///         p.absolutize_virtually("/virtual/root").unwrap_err().kind()
    ///     );
    /// }
    /// ```
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::path::Path;
    ///
    /// use std::io::ErrorKind;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("../path/to/123/456");
    ///
    ///     assert_eq!(
    ///         ErrorKind::InvalidInput,
    ///         p.absolutize_virtually("/virtual/root").unwrap_err().kind()
    ///     );
    /// }
    /// ```
    ///
    /// A path which does not start with a `MAIN_SEPARATOR`, **Single Dot** and **Double Dots**, will be located in the virtual root after the `absolutize_virtually` method is used.
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::path::Path;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("path/to/123/456");
    ///
    ///     assert_eq!(
    ///         "/virtual/root/path/to/123/456",
    ///         p.absolutize_virtually("/virtual/root").unwrap().to_str().unwrap()
    ///     );
    /// }
    /// ```
    ///
    /// ```
    /// extern crate path_absolutize;
    ///
    /// use std::path::Path;
    ///
    /// use path_absolutize::*;
    ///
    /// if cfg!(not(windows)) {
    ///     let p = Path::new("path/to/../../../../123/456");
    ///
    ///     assert_eq!(
    ///         "/virtual/root/123/456",
    ///         p.absolutize_virtually("/virtual/root").unwrap().to_str().unwrap()
    ///     );
    /// }
    /// ```
    fn absolutize_virtually<P: AsRef<Path>>(&self, virtual_root: P) -> io::Result<PathBuf>;
}

impl Absolutize for PathBuf {
    fn absolutize(&self) -> io::Result<PathBuf> {
        let path = Path::new(&self);

        path.absolutize()
    }

    fn absolutize_virtually<P: AsRef<Path>>(&self, virtual_root: P) -> io::Result<PathBuf> {
        let path = Path::new(&self);

        path.absolutize_virtually(virtual_root)
    }
}

mod unix;
mod windows;