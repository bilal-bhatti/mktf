# Make Tofu

If you have OpenTofu files in JSON format, which are easier to process using existing JSON tools, and want to convert them to HCL format, this is a simple command line tool to do just that.

``` sh
go install github.com/bilal-bhatti/mktf@v0.0.1
```

## Usage
* `mktf`
    * convert all *.tf.json files to *.tf in the current directory
* `mktf -`
    * read from stdin and write to stdout
* `mktf one.tf.json two.tf.json`
    * convert provided list of *.tf.json files
