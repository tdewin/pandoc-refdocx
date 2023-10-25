# How to use

```bash
pandoc -o test.docx --reference-doc=reference.docx test.md
```


# Deriving from
```bash
BEFORE=$(pwd)
mkdir -p src && cd src
pandoc -o reference.docx --print-default-data-file reference.docx
unzip reference.docx  
unzip -l reference.docx | tr -s " " | cut -d" " -f 5 | grep -v "^$" | grep  -v "^Name$" > ziplist.txt

echo "[" > ziplist.json
FIRST=1
while read p; do
  if [ $FIRST = "1" ] ; then FIRST=2; else ; echo "," >>ziplist.json;fi
  echo -n "\"$p\"">> ziplist.json
done <ziplist.txt
echo "]" >> ziplist.json

cat ziplist.json | jq

cd $BEFORE
```

zipjs
```bash
BEFORE=$(pwd)
mkdir -p zipjs && cd zipjs
git clone https://github.com/gildas-lormeau/zip.js.git

```

upload to a server or use go to run a static server
```bash
go run staticserver.go 
```

