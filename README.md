# How to use
Live : [https://tdewin.github.io/pandoc-refdocx/](https://tdewin.github.io/pandoc-refdocx/)

```bash
pandoc -o test.docx --reference-doc=reference.docx test.md
```

# Extracting from a document you already have and wanting to get reffed
requires xmlstarlet; but more allows for searching like vi (eg esc+/+pattern)
```bash
DOCXREF=reference.docx
unzip -p $DOCXREF word/styles.xml | xmlstarlet fo | more
unzip -p $DOCXREF word/document.xml | xmlstarlet fo | more
unzip -p $DOCXREF word/theme/theme1.xml | xmlstarlet fo | more
```

even more advanced, get the sz in 20*pts
```bash
unzip -p $DOCXREF word/document.xml | xmlstarlet sel -t -c  "/w:document/w:body/w:sectPr"  | xmlstarlet fo
```

or 
```bash
DOCXML=$(unzip -p $DOCXREF word/document.xml)
function getPath {
  printf "%s: %s\n" $1 $(echo $DOCXML | xmlstarlet sel -t -v  $1)
}
getPath "/w:document/w:body/w:sectPr/w:pgSz/@w:w"
getPath "/w:document/w:body/w:sectPr/w:pgSz/@w:h"
getPath "/w:document/w:body/w:sectPr/w:pgSz/@w:orient"
getPath "/w:document/w:body/w:sectPr/w:pgMar/@w:top"
getPath "/w:document/w:body/w:sectPr/w:pgMar/@w:bottom"
getPath "/w:document/w:body/w:sectPr/w:pgMar/@w:left"
getPath "/w:document/w:body/w:sectPr/w:pgMar/@w:right"
getPath "/w:document/w:body/w:sectPr/w:pgMar/@w:header"
getPath "/w:document/w:body/w:sectPr/w:pgMar/@w:footer"
```



# Dev notes
## Deriving from
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


