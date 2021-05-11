BASE := devbase

REPO := gitrhythm/"${BASE}"
TAG := latest
CONTAINER_NAME := "${BASE}"

IMAGE := $(shell docker image ls ${REPO} --format {{.Repository}})
CONTAINER := $(shell docker container ps -a --format {{.Names}} | grep ${CONTAINER_NAME})
UP_CONTAINER := $(shell docker container ps --format {{.Names}} | grep ${CONTAINER_NAME})

b:	# イメージのビルド(build)
	@if [[ "${IMAGE}" == "" ]]; then \
		echo イメージをビルドしています... ;\
		docker build -t ${REPO}:${TAG} . ;\
	fi

c:	# イメージを削除
	@make s
	@if [[ "${CONTAINER}" == "${CONTAINER_NAME}" ]]; then \
		echo コンテナを削除しています... ;\
		docker container rm "${CONTAINER_NAME}" ;\
	fi
	@if [[ "${IMAGE}" == "${REPO}" ]]; then \
		echo イメージを削除しています ;\
		docker image rm ${REPO} ;\
	fi

r:	# コンテナの作成・起動(run)
	@make b
	@if [[ "${CONTAINER}" == "" ]]; then \
		echo コンテナを作成しています... ;\
		docker container create -it --name ${CONTAINER_NAME} ${REPO}:${TAG} ;\
	fi
	@if [[ "${UP_CONTAINER}" == "" ]]; then \
		echo コンテナを起動しています... ;\
		docker container start ${CONTAINER_NAME} ;\
	fi

s:	# コンテナ停止(stop)
	@if [[ "${UP_CONTAINER}" == "${CONTAINER_NAME}" ]]; then \
		echo コンテナを停止しています... ;\
		docker container stop "${CONTAINER_NAME}" ;\
	fi

l:	# コンテナに接続(ログイン)(login)
	@make b
	@make r
	docker exec -it ${CONTAINER_NAME} /bin/bash --login

cb: # イメージの作り直し
	@make c
	@make b

cl:	# イメージを作り直してログイン
	@make c
	@make l

h:
	@echo "b: イメージのビルド(build)"
	@echo "c: イメージの削除(clean) 〜 コンテナが起動・作成されていれば停止・削除する"
	@echo "r: コンテナの作成(run)、及び起動"
	@echo "s: コンテナの停止(stop)"
	@echo "l: コンテナへのログイン(login)"
	@echo "cb: イメージの作り直し"
	@echo "cl: イメージを作り直し 〜 コンテナ作成・起動 〜 ログイン"
	@echo "echo: 定義済み変数の表示"

echo:
	@echo "REPO:${REPO}"
	@echo "TAG:${TAG}"
	@echo "CONTAINER_NAME:${CONTAINER_NAME}"
	@echo "IMAGE:${IMAGE}"
	@echo "CONTAINER:${CONTAINER}"
	@echo "UP_CONTAINER:${UP_CONTAINER}"
