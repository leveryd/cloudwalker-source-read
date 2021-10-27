# 竞品分析，其他厂商是否可以识别样本


function changting(){
  r=$(curl https://webshellchop.chaitin.cn -F "initialPreviewConfig=[]" -F "initialPreview=[]" -F "initialPreviewThumbTags=[]" -F "fileId=71_eval.php" -F "inputfile=@$1" -x "http://127.0.0.1:8081" 2>/dev/null | jq .data.detail)
  if [ "$r" == "{}" ];then
    echo "长亭：normal"
  else
    echo "长亭：evil"
  fi
}

function webdir(){
  md5=$(curl https://scanner.baidu.com/enqueue -F "archive=@$1" -x "http://127.0.0.1:8081" 2>/dev/null | jq .md5)
  echo $md5
  curl  "https://scanner.baidu.com/result/$md5"
}

function cdxy(){
  result=$(curl http://webshell.cdxy.me/api -F "file=@$1" -x "http://127.0.0.1:8081" 2>/dev/null | grep -i "danger")
  if [ -z $result ];then
    echo "cdxy：normal"
  else
    echo "cdxy：evil"
  fi
}

function handle(){
  path=$(cat $1 | awk '{print $3}')
  for i in $path;do
    echo $i;
    changting $i
    sleep 1
  done
}

# webdir /tmp/php/1.php
handle /tmp/x.html
