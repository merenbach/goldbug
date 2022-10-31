
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

* Explore making MASC String() into Printable()
* Fix test names and filename conventions
* Implement PASC transforms and tests
* Refactor testing methodology for MASC and PASC to reduce duplication
* Explore alternative PASC with (slope, intercept) = (-1, -1)
* Expand tests for custom ciphertext alphabets
* Improve and enhance test cases overall
* Expand tests for tabula recta to catch cases with reversed or alternate alphabets
* Allow grouping/delimiters
* Make grids one-based instead of zero-based in some contexts? (e.g., Polybius)
* Allow transposition cipher key to be for columns directly, rather than relative ordering
* Allow number duplicates to be ordered backwards when creating lexicographic key
* Add more to-string printable options to transposition and other ciphers
* Add printout option to columnar transposition ciphers--and use tab writer
* Include primer at beginning of GROMARK text and end with check digit per <https://www.cryptogram.org/downloads/aca.info/ciphers/Gromark.pdf>?
* Add null options to transposition cipher? Special null adder utility module?
* Add more MASC variant caseless tests
* Add shift and directionality to Trithemius?
* Nulls for rail fence?
* Improve autoclave further
* Add keyed vigenere (Quagmire IV)
