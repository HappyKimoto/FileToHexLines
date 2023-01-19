# file to hex lines

## Objective
- Analyze text files with several character encodings in Hex.
- Compare line by line between raw text file agaisnt hex string.

## How to use
- Pass a folder or a file as the first argument to filetohexlines.exe.
    - Case1: Run filetohexlines.exe from command line and drag/drop the file or the folder.
    - Case2: Slace filetohexlines.exe to SendTo folder and execute from right click on file explorer.
- Converted text file will be saved on the same directory.

## Usage
- UTF16 produces inaccurate result because the new line is **\r\0\n\0**.
```
C:\...>filetohexlines.exe .\sample
=== File To Hex Lines ===
line = 3
270 bytes written to sample\sample-ShiftJIS.txt.FileToHexLines.txt
line = 4
323 bytes written to sample\sample-UTF16.txt.FileToHexLines.txt
line = 3
393 bytes written to sample\sample-UTF8.txt.FileToHexLines.txt

C:\....\FileToHexLines>
```