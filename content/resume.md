# Reed Allman

<rdallman10@gmail.com> · [rdallman.me]

-------------------------------------------------------------------

I'm currently building [usually distributed] systems, but am interested in most
things.

### Labor

**May 2017-present, Principal Member of Technical Staff, Oracle**

  *  Core contributor on [fn], an open source Functions as a Service (FaaS)
     product that leverages Docker to run any language in a FaaS fashion.
     Primarily focused on load balancing, scalability, task execution, storage
     and messaging.
  *  Building out a hosted version of [fn] on Oracle's cloud.

**May 2014-April 2017, Senior Backend Engineer, [Iron.io]**

  *  Co-built a strongly consistent, distributed key-value store on top of RocksDB.
     Built distributed transactions, node membership, dynamic load balancing,
     automagic sharding.
  *  Co-built a distributed message queue system on top of that, specifically
     for performance + persistence for processing [delayed] jobs. [It scales]
  *  Built authentication service on top of said distributed key-value store,
     as well (much easier than the queue).
  *  Increased cluster utilization from 40% to 80% by building a custom
     autoscaler for job processing servers. Saved lots of $$$ and zzz,
     didn't have to launch servers by hand anymore.
  *  Migrated job runner from Ruby to Go. Decreased p99 task start time 100x,
     p75 by 30x by simplifying API and intelligent, probabilistic queue
     polling. Eliminated issues of jobs getting stuck in queue.
  *  Migrated infrastructure from upstart scripts and binaries to CoreOS,
     systemd and Docker. Setup push button releases for all products.
  *  Wrestled with Linux. Battled Docker. Wrote a lot of Go code. Ran a lot of
     Go code.

**May 2013-May 2014, Undergraduate Research Assistant, Auburn University** 

  *  Built refactoring tools. Funded by Google to work on refactoring Go. C:
     [OpenRefactory] Go: [godoctor]
  *  Built a statement level control flow graph for Go source code and used
     that to do data flow analyses.
  *  Constructed a pretty awesome testing infrastructure to test our tool on
     all the Go source I could find (4.5M lines).
  *  Developed the CLI, Sublime Text plugin, Vim plugin and JSON protocol for
     C and Go refactoring tools.

### Education

**B.S., Software Engineering, Auburn University, 2014.**

  *  Minor in Business-Engineering-Technology ([B-E-T] TL;DR engineering entrepreneurship)

### Contributions

  *  I dabble in open source: [github]
  *  Sometimes I write: [medium]
  *  Occasionally I can be convinced to stand in front of people: [speakerdeck]
  *  Co-organizer of [GoSF] and [RocksDB] meetups

[It scales]:https://www.iron.io/1m-msgsec-ironmqv3-hits-dos-commas/
[B-E-T]:https://eng.auburn.edu/research/centers/twc/bet-program/index.html
[github]:https://github.com/rdallman
[godoctor]:https://github.com/godoctor/godoctor
[Iron.io]:https://iron.io
[OpenRefactory]:https://dl.acm.org/citation.cfm?id=2541349
[rdallman.me]:https://www.rdallman.me
[speakerdeck]:https://speakerdeck.com/rdallman
[GoSF]:https://www.meetup.com/golangsf
[RocksDB]:https://www.meetup.com/RocksDB
[medium]: https://medium.com/@rdallman10
[fn]:https://github.com/fnproject/fn
