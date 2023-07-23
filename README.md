# resolvconf-reset

This simple Go program was built to resolve my hardware issues on Manjaro's network driver. Sometimes when I update network drivers or utilities, the `name_servers` DNS used by `pacman` to check internet availability are reseted to some DNS that are not available. So running this program as super user makes updating pacman's dns more convenient, for example:

```bash
sudo resolvconf-reset 8.8.8.8 8.8.4.4 # Google's DNS but can be any of your choice
```

## How it works

Following [Manjaro Networking Guide](https://wiki.manjaro.org/index.php/Networking), the arguments passed to program will override `name_servers` field at `/etc/resolvconf.conf` file, that holds which DNS can be used to check your network before download any package. After overriding the file, will be executed `resolvconf -u` to update `/etc/resolv.conf` file.
