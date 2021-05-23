**I would recommend using [NtHiM](https://github.com/TheBinitGhimire/NtHiM/tree/main/src)**
# takeover

Since [Ice3man543](https://github.com/Ice3man543) decided to discontinue the devlopment of their tool i.e SubOver, I thought it would be nice to revive this tool.

## FAQs

__Why?__

The major reason behind reviving this tool would be to be able to perform subdomain takeover check on a mass scale. By mass scale I mean is that sometime what happens in that you have a program with a very large scope and loads of domains. Now if you're a monster than you'll have all the subdoma under a single file called `domain.txt` or something similar. But if you are a sane person than you'd like to keep all subdomains of a root domain in one file and this way there can be loads of files. So it would be better if we can just pass the path to the folder and relax.

__Why not use one liner?__

yeah I know that I can write a one liner combined with tools like nuclei or subjack but I kind don't want to :)

__Why not use subjack?__

I don't know why but I've had issue in installing subjack on digital ocean VPN. Whenever I try to pull it using `go get` it just hangs there. Also I am not sure if its under active development cause I noticed there are loads of pending issues and Pull requests.

__Can I use different Providers list?__

Currently no, but I plan to add this feature so like you can use file from subjack or if you make your own.


## Options

```
  -d string
        directory having files of domains
  -https
        Force HTTPS connections
  -l string
        List of hosts to check takeovers on
  -p string
        Path of the providers file
  -t int
        Number of threads to use (default 20)
  -timeout int
        Seconds to wait before timeout (default 10)
  -v    Show verbose output
```

## Usage

The usage is same as it was.

* For hunting on the same list:

```
takeover -l <subdomain-list.txt>
```

* For testing a directory with all the subdomain list

```
takeover -d <directory>
```

Make sure the directory doesn't have any other file otherwise the process will take extra time. Also the tool might crash(I'm still learning golang.)

## Installation

You can download the binary from the release page. Also if you want you can clone this repository and build the binary yourself.

If you have go compiler installed then you can use go get github.com/mzfr/takeover.

__NOTE__: takeover uses `provider.json` file. So either have a file named providers.json in your current working directory  you can provide the path via `-p` flag.

## Acknowledgements and Credits

Thanks to [Ice3man543](https://github.com/Ice3man543) for making [SubOver](https://github.com/Ice3man543/SubOver)

# Support

If you'd like you can buy me some coffee:

<a href="https://www.buymeacoffee.com/mzfr" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" style="height: 51px !important;width: 217px !important;" ></a>
