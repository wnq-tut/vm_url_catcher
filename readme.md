# vm_url_catcher 

A small tool running in a virtual machine (via RemoteApp) to quickly copy URL clicks to the clipboard.  

Self use. 自用. 

## Usage

Build:

```bash
go build -ldflags="-H windowsgui" -o urlcatcher.exe
```

Import `register-urlcatcher.reg` to register this tool as the handler for HTTP and HTTPS protocols.  
Before importing, modify the path inside register-urlcatcher.reg:

```bash
"C:\\***path***\\urlcatcher.exe\"
```

Replace with the actual absolute path where `urlcatcher.exe` is located.  

Then open **Settings → Apps → Default apps → Choose defaults by protocol**,  
find the `HTTP` and `HTTPS` protocols, and set them to open with `urlcatcher.exe`.

This will cause all HTTP/HTTPS links to launch this tool, passing the URL as an argument.

## License

MIT