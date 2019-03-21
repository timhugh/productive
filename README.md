# Productive

Take care of those distractions!

## Usage

I make no claims that this will work on anything other than MacOS (although I'm sure Linux isn't far off).
Also, I have static IPs for my docker machines. More on that here: https://github.com/timhugh/dotfiles/blob/0192931798517f7e991ad284b5dbf2c077429faf/bash/bash_profile.d/docker.sh

```
docker build -t productive:latest .
docker run -d --rm -p "80:80" productive
```

Add to your /etc/hosts file:
```
192.168.99.101 facebook.com
192.168.99.101 login.facebook.com
192.168.99.101 www.facebook.com
```

Flush DNS cache:
```
dscacheutil -flushcache
```

Et voila!

```
~ $ http GET www.facebook.com
HTTP/1.1 200 OK
Content-Length: 5
Content-Type: text/plain; charset=utf-8
Date: Thu, 21 Mar 2019 17:12:04 GMT

Nope!
```
