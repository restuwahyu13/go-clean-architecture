
DB_DIR := $(realpath ./)
SQLC := $(realpath ${DB_DIR}/node_modules/.bin/sequelize-cli)
BUILD := npm run build

##################################
# Database Migration Territory
##################################
sqlc:
	${SQLC} -h

mig-build:
	rm -r ${DB_DIR}/dist; ${BUILD}

mig-create:
ifdef f
	${SQLC} migration:create --name ${f}
endif

mig-rollback:
ifdef f
	${SQLC} db:migrate:undo --name ${f}
endif

mig-status:
	${SQLC} db:migrate:status --debug

mig-upf:
ifdef f
	rm -r ${DB_DIR}/dist; ${BUILD}
	${SQLC} db:migrate --name ${DB_DIR}/dist/migrations/${f}
endif

mig-up:
	rm -r ${DB_DIR}/dist; ${BUILD}
	${SQLC} db:migrate

mig-down: # DANGER COMMAND, BE CAREFUL CAN DELETE ALL TABLES
	${SQLC} db:migrate:undo:all

seed-create:
ifdef f
	${SQLC} seed:create --name ${f}
endif

seed-upf:
ifdef f
	rm -r ${DB_DIR}/dist; ${BUILD}
	${SQLC} db:seed --seed ${DB_DIR}/dist/seeders/${f}
endif

seed-up:
	rm -r ${DB_DIR}/dist; ${BUILD}
	${SQLC} db:seed:all

seed-down:
	rm -r ${DB_DIR}/dist; ${BUILD}
	${SQLC} db:seed:undo:all