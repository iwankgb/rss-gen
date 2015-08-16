[![Build Status](https://travis-ci.org/iwankgb/rss-gen.svg?branch=master)](https://travis-ci.org/iwankgb/rss-gen)

rss-gen
=======

This is just a simple tool that allows to build RSS xml files. See [example.yaml](resources/example.yaml) for self explanatory yaml input. I know that parsing yaml to xml might look silly to some of you but I hate idea of maintaining XML file manually...

I do use this tools - see [subjective news review](http://critical.today/files/subiektywny.xml) (Polish only as this is not intended for international audience).

Command line options
--------------------

* ``-yaml`` - path to input yaml files
* ``-rss-existing`` - path to existing RSS file (useful to avoid updating existing items' dates)
* ``-count`` - Number of items to be included in the feed
