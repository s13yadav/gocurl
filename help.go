package main

func printHelp(){

    help:=`*************  gocurl help  *************

gocurl [-X|-d|--connect-timeout|-retry|-retry-delay] -u <URL>

      --connect-timeout <seconds>
              Maximum time in seconds that you allow the connection to the server to take.  This only limits the connection phase, once curl has connected this option is of no more use.

      -d <data>
              (HTTP) Sends the specified data in a POST request to the HTTP server, in the same way that a browser does when a user has filled in an HTML form and presses the submit button. This will cause curl to pass
              the data to the server using the content-type application/x-www-form-urlencoded.  Compare to -F, --form.

              -d, --data is the same as --data-ascii. To post data purely binary, you should instead use the --data-binary option. To URL-encode the value of a form field you may use --data-urlencode.

              If  any  of  these  options is used more than once on the same command line, the data pieces specified will be merged together with a separating &-symbol. Thus, using '-d name=daniel -d skill=lousy' would
              generate a post chunk that looks like 'name=daniel&skill=lousy'.

              If you start the data with the letter @, the rest should be a file name to read the data from, or - if you want curl to read the data from stdin.  The contents of the file  must  already  be  URL-encoded.
              Multiple files can also be specified. Posting data from a file named 'foobar' would thus be done with --data @foobar.


      --retry <num>
              If a transient error is returned when scurl tries to perform a transfer, it will retry this number of times before giving up. Setting the number to 0 makes curl do no retries (which is the default).

      --retry-delay <seconds>
              scurl sleep this amount of time before each retry when a transfer has failed with a transient error. This option is only interesting  if --retry is also used.

      -X
              (HTTP) Specifies a custom request method to use when communicating with the HTTP server.  The specified request will be used instead of the method otherwise used (which defaults to GET).


Example:
        scurl -X POST --connect-timeout 120 --retry 5 --retry-delay 5 -d "{Key1:Value1}" https://10.255.14.111:5572/api/v1/executecmd

    `
    println(help)
}

