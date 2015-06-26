@echo off
cargo clean
cargo update
rustc --version
set stime=%time%
echo (%stime%) build start..
cargo build --release
set etime=%time%
echo (%etime%) build finish..
rustc --version
echo start:(%stime%)
echo finish:(%etime%)
pause
pause