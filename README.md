# deblocus

Current version is v0.11.2770-beta

[![Build Status](https://travis-ci.org/Lafeng/deblocus.svg?branch=master)](https://travis-ci.org/Lafeng/deblocus)
[![MIT License](https://img.shields.io/packagist/l/doctrine/orm.svg)](http://opensource.org/licenses/MIT)
[![Issues](https://img.shields.io/github/issues/Lafeng/deblocus.svg)](https://github.com/Lafeng/deblocus/issues)
[![Comment-on-Github.Party](https://img.shields.io/badge/Comment%20on-Github.Party-yellow.svg)](https://github.party/item?id=46)

## Introduction

deblocus is a tunneling software that allows secure TCP connection through the server by acting as a SOCKS5/HTTP proxy server on the client side.

## Features

- Traffic encrypted with dynamic session key
	* deblocus has [perfect forward secrecy](https://en.wikipedia.org/wiki/Forward_secrecy).
- Tunnel's communication was provided with mildly obfuscated traffic based on special protocol
	* Keep communications secure and privacy over untrusted networks.
- TCP multiplexing inside tunnel and Fast-Open over application layer
	* Optimized connectivity for massive short requests, latency greatly reduced to nearly 1-RTT.
- Filtering traffic to specified country (using GeoLite2 data)
	* You can avoid the proxy was used for some inappropriate purposes.
- High-performance encryption implementation
	* Could enjoy high throughput and low resource consumption.

## Quickstart

Available on AMD64 of Windows, Linux and OS X.

1. Download [executables](https://github.com/Lafeng/deblocus/releases) or compile it from source.
2. Read [Wiki](https://github.com/Lafeng/deblocus/wiki) to learn more.

## Acknowledgements

Thanks to these projects: [qtunnel](https://github.com/getqujing/qtunnel), [osext](https://bitbucket.org/kardianos/osext), [dhkx](https://github.com/monnand/dhkx), [glog](https://github.com/golang/glog), [siphash](https://github.com/dchest/siphash) and [chacha-opt](https://github.com/floodyberry/chacha-opt).

This product includes GeoLite2 data created by MaxMind, available from <http://www.maxmind.com>.

## License

[MIT License](https://github.com/Lafeng/deblocus/blob/master/LICENSE)


# deblocus

## 介绍

安全高效的局部代理(Socks5/HTTP)工具，可满足在企业网络/公共WiFi等场景下安全通信的需要。

## 使用

首先阅读 [Wiki](https://github.com/Lafeng/deblocus/wiki)，从项目Releases下载程序，或自行编译。
