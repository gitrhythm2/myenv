#!/usr/bin/env fish

# fish用組み込み関数サンプル
function o
  set -l item (./opener.sh $argv[1] | string split " ")
  if test (count $item) -lt 2
    exit 1
  end

  switch $item[1]
  case file
    vi $item[2]
  case dir
    cd $item[2]
    ls -laF
  case www
    open -a Google\ Chrome https://"$item[2]"
  case bear
    open "bear://x-callback-url/open-note?id=$item[2]"
  end
end

o $argv
