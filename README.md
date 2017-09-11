# DMARC Parser

Parses Google DMARC report and reports if some tests not passed.

## Build

```
go install
```

## Usage
Read from standard input:

```
zcat google.com!example.com!123456786543!1234567876543.zip | dmarc-report
```

Read from file:

```
dmarc-report -fileName google.com!example.com!123456786543!1234567876543.xml
```
