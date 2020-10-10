# SubOver 2.0

Since [Ice3man543](https://github.com/Ice3man543) decided to discontinue the dvelopment of their tool i.e SubOver, I thought it would be nice to revive this tool.

## Summary

__Why?__

The major reason behind reviving this tool would be to be able to perform subdomain takeover check on a mass scale. By mass scale I mean is that sometime what happens in that you have a program with a very large scope and loads of domains. Now if you're a monster than you'll have all the subdoma under a single file called `domain.txt` or something similar. But if you are a sane person than you'd like to keep all subdomains of a root domain in one file and this way there can be loads of files. So it would be better if we can just pass the path to the folder and relax.

__Why not use one liner?__

yeah I know that I can write a one liner combined with tools like nuclei or subjack but I kind don't want to :)


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

Make sure the directory doesn't have any other file otherwise the process will take extra time.
