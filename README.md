# jmeter_hello_world
Hello world for jmeter

A very basic example of Jmeter test plan based on (http://jmeter.apache.org/usermanual/build-web-test-plan.html)

## Installation

### Jmeter

    sudo apt-get install jmeter jmeter-http

### Ruby API requirements

    sudo apt-get install ruby
    sudo gem install sinatra

### Go API requirements

    sudo apt-get install golang
    cd jmeter_hello_world
    gobuild api

Based on (http://soryy.com/blog/2014/not-another-go-net-http-tutorial/)

## Running JMeter tests

Fire up an API.  Both examples run on port 8080 so you can only run one at a time.
Both API implementations are equivalent providing a `\random` endpoint which takes an optional `size` parameter.
If the size parameter is missing a random number of elements will be returned 0 < x < 200

1. Ruby api

    `ruby api.rb`

2. Go

    `./api`

To test your api is running, in another shell run

    curl localhost:8080/random?size=2

### Jmeter

Start JMeter

    jmeter

Once Jmeter has fired up, load up the test plan (File -> Open) and select 'SimpleTestPlan.jmx'
To run the test do 'Run -> Start' and navigate to the 'Graph Results' in the tree on the left to observe the test performance.

## Credits

Inspired by (https://github.com/mjgpy3/jmeter_toy)