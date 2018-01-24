# Use goreman (`go get -u github.com/mattn/goreman`) to run
#   goreman -f Procfile start
#
# This bootstraps a new cluster. To join/restart remove
# the `--cluster-bootstrap` flag
#
# https://devcenter.heroku.com/articles/procfile
#
tweetBank1: tweetBank start cluster1 --id 1 --cluster-port=4747 --cluster-bootstrap
tweetBank2: tweetBank start cluster1 --id 2 --cluster-port=4848
tweetBank3: tweetBank start cluster1 --id 3 --cluster-port=4949
