#!/usr/bin/env bash

echo -e "CLI config-profile:"
echo -e "Default profile"
./cli.sh config-profile set FOO=DEFAULT_BAR
./test/assert "should apply config" \
$(./cli.sh checkenv | grep "FOO") "FOO=DEFAULT_BAR"

./cli.sh cleanconf
./test/assert "should clean config" \
$(./cli.sh checkenv | grep "FOO") ""

./cli.sh config-profile set FOO=DEFAULT_BAR_CHANGED
./test/assert "should change config" \
$(./cli.sh checkenv | grep "FOO") "FOO=DEFAULT_BAR_CHANGED"

echo -e "<ANY> Profile:"
./cli.sh config-profile -n test set FOO=BAR
./test/assert "should apply config" \
$(./cli.sh --profile test checkenv | grep "FOO") "FOO=BAR"

./test/assert "should not affect on Default profile config" \
$(./cli.sh checkenv | grep "FOO") "FOO=DEFAULT_BAR_CHANGED"

./cli.sh config-profile -n test set FOO=BAR_CHANGED
./test/assert "should change config" \
$(./cli.sh --profile test checkenv | grep "FOO") "FOO=BAR_CHANGED"

./cli.sh config-profile -n test clean
./test/assert "should clean config" \
$(./cli.sh --profile test checkenv | grep "FOO") ""

./cli.sh -p var set THIS=THIS SHOULD=SHOULD BE=BE CLI=CLI VARIABLES=VARIABLES
./test/assert "should apply cli env variables" \
$(./cli.sh -p var run ./test/run_test.sh | grep "THIS") "THISSHOULDBECLIVARIABLES"
