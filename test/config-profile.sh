#!/usr/bin/env bash

echo -e "CLI config-profile:"
echo -e "Default profile"
./cli.sh config-profile set FOO=DEFAULT_BAR
./test/assert "should apply config" \
$(./cli.sh checkconf | grep "FOO") "FOO=DEFAULT_BAR"

./cli.sh config-profile set FOO=DEFAULT_BAR_CHANGED
./test/assert "should change config" \
$(./cli.sh checkconf | grep "FOO") "FOO=DEFAULT_BAR_CHANGED"

echo -e "<ANY> Profile:"
./cli.sh config-profile -n test set FOO=BAR
./test/assert "should apply config" \
$(./cli.sh --profile test checkconf | grep "FOO") "FOO=BAR"

./test/assert.sh "should not affect on Default profile config" \
$(./cli.sh checkconf | grep "FOO") "FOO=DEFAULT_BAR_CHANGED"

./cli.sh config-profile -n test set FOO=BAR_CHANGED
./test/assert "should change config" \
$(./cli.sh --profile test checkconf | grep "FOO") "FOO=BAR_CHANGED"
