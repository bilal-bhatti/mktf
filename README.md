# Make Tofu

If you have OpenTofu in json format, which are much easier to process using existing JSON tools, and want to convert them to HCL format, this is a simple command line tool to make that easier.

``` sh
go install https://github.com/bilal-bhatti/mktf
```

## Usage
* mktf
 * convert all *.tf.json file to *.tf
* mktf -
 * read from stdin out write to stdout
* mktf [list of *.tf.json files]
