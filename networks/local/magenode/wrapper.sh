#!/usr/bin/env sh

BINARY=/mgd/linux/${BINARY:-mgd}
echo "binary: ${BINARY}"
ID=${ID:-0}
LOG=${LOG:-mgd.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'mgd' E.g.: -e BINARY=mgd_my_test_version"
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export MGDHOME="/mgd/node${ID}/mgd"

if [ -d "$(dirname "${MGDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${MGDHOME}" "$@" | tee "${MGDHOME}/${LOG}"
else
  "${BINARY}" --home "${MGDHOME}" "$@"
fi
