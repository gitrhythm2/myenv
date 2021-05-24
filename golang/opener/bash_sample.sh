#!/usr/bin/env bash

# bash用組み込み関数サンプル
o() {
  declare -a item=($(opener $1))
  if [[ ${#item[@]} -lt 2 ]]; then
    return
  fi

  case "${item[0]}" in
  file)
    vi "${item[1]}" ;;
  dir)
    cd "${item[1]}"
    ls -laF
    ;;
  www)
    open -a Google\ Chrome https://"${item[1]}" ;;
  bear)
    open bear://x-callback-url/open-note?id="${item[1]}" ;;
  esac
}

o $1
