
# Gold-Bug

A Web app written in Go to perform encipherment and decipherment of messages using old-fashioned field ciphers.

## Configuration

Instantiate your deployment pipeline as follows, adapting as necessary the parameter overrides (including any not shown here):

    aws cloudformation deploy \
        --template-file pipeline.yaml \
        --stack-name pingaling-dev \
        --capabilities CAPABILITY_IAM \
        --parameter-overrides \
            BranchName=develop \
            RepositoryName=gold-bug \
            GitHubOwner=merenbach \
            AlarmEmail='your_name@your_domain.com' \
            GitHubOAuthToken='YOUR_PAT_HERE'

 The first deployment should trigger automatically. If it does not, simply push to the branch to kick off the deployment process.

## TODO

* MASC casing preservation/normalization?
* Add shift and directionality to Trithemius?
* Nulls for rail fence?
* Add keyed vigenere (Quagmire IV)
