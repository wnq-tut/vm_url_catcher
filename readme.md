# vm_url_catcher 

A small tool running in a virtual machine (via RemoteApp) to quickly copy URL clicks to the clipboard.  

Self use. 自用. 

## Usage

Build:

```bash
go build -ldflags="-H windowsgui" -o urlcatcher.exe
```
Import `register-urlcatcher.reg`.  

Then open **Settings → Apps → Default apps → Choose defaults by protocol**,  
find the `HTTP` and `HTTPS` protocols, and set them to open with `urlcatcher.exe`.

This will cause all HTTP/HTTPS links to launch this tool, passing the URL as an argument.

## License

MIT