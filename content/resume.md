# Reed Allman

<reed@rdallman.com> · [rdallman.com]

-------------------------------------------------------------------

I'm currently building distributed compute & storage systems [as-a-Service],
but am interested in many things.

### Labor

**May 2017-present, Principal Member of Technical Staff, Oracle**

  *  Core contributor on [Fn], an open source Functions as a Service (FaaS)
     product that leverages Docker to run any container in a FaaS fashion.
     Primarily focused on load balancing, scalability, task execution, storage
     and messaging.
  *  Involved in building out a hosted version of [Fn] on Oracle's cloud.
  *  Re-architected task running agent in order to reduce start latencies,
     making it possible for tasks to wait on existing containers of a function
     kind to become available to run while a new container may be spawned,
     where previously most of those tasks would time out. This also brought
     down the number of containers launched during spikes by a painfully
     large, unknown factor (early days!).
  *  Led API 2.0 redesign effort, after spending many months working with our
     v1 API we discovered it was unintuitive when adding hooks into various
     services to trigger tasks: an API gateway, an s3-like store, etc. This
     led to an effort to rework our API to allow functions to be their own
     entity, and to have a one-to-many relationship with triggers (as is
     industry-standard for FaaS).
  *  Led effort to separate API nodes from task running nodes. We needed a way
     to separate these in order to put task running nodes in a separate network
     from API nodes, where API nodes are a touch point for users to deploy and
     manage functions on, but task running nodes operate on a separate network.
  *  Led effort to instrument all code with [OpenCensus], a distributed
     tracing and statistics library with plugs for many backends. Used these
     to reduce latency in our critical path more than once; reduced
     allocations, removed docker logging, batched writes, among others.
  *  Presented talks at Devoxx, GoSF, and various meetups about Fn. Wrote a
     number of blog posts for [Fn medium]. Helped users debug their computers
     in Slack. Reviewed a zillion Pull Requests. Wore a T-shirt bearing the
     Fn logo once a week around SF to fulfill Peter Thiel's prophecy.

**May 2014-April 2017, Senior Backend Engineer, [Iron.io]**

  *  Co-built a strongly consistent, distributed key-value store on top of RocksDB.
     Built distributed transactions, node membership, dynamic load balancing,
     automagic sharding, and automagic rebalancing.
  *  Co-built a distributed message queue system on top of that, specifically
     for performance + persistence for processing [delayed] jobs. [It scales]
  *  Built authentication service on top of said distributed key-value store,
     as well (much easier than the queue).
  *  Increased cluster utilization from <40% to >80% by building a custom
     autoscaler for job processing servers. Saved lots of $$$ and zzz,
     didn't have to launch servers by hand anymore.
  *  Migrated job runner from Ruby to Go. Decreased p99 task start time 100x,
     p75 by 30x by simplifying API and intelligent, probabilistic queue
     polling. Eliminated issues of jobs getting stuck in queue.
  *  Migrated infrastructure from upstart scripts and binaries to CoreOS,
     systemd and Docker. Setup push button releases for all products.
  *  Wrestled with Linux. Battled Docker. Wrote a lot of Go code. Ran a lot of
     Go code. Contributed to RocksDB, CoreOS. Presented talks at meetups and
     conferences about IronMQ. Wrote blog posts.

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

**Bachelor of Software Engineering, Auburn University, 2010-2014**

  *  Minor in Business-Engineering-Technology ([B-E-T] TL;DR engineering entrepreneurship)

### Contributions

  *  I dabble in open source: [github]
  *  Sometimes I write: [medium]
  *  Occasionally I can be convinced to stand in front of people: [speakerdeck]
  *  Former co-organizer of [GoSF] and [RocksDB] meetups

[It scales]:https://www.iron.io/1m-msgsec-ironmqv3-hits-dos-commas/
[B-E-T]:https://eng.auburn.edu/research/centers/twc/bet-program/index.html
[github]:https://github.com/rdallman
[godoctor]:https://github.com/godoctor/godoctor
[Iron.io]:https://iron.io
[OpenRefactory]:https://dl.acm.org/citation.cfm?id=2541349
[rdallman.com]:https://rdallman.com
[speakerdeck]:https://speakerdeck.com/rdallman
[GoSF]:https://www.meetup.com/golangsf
[RocksDB]:https://www.meetup.com/RocksDB
[medium]: https://medium.com/@rdallman10
[Fn]:https://github.com/fnproject/fn
[OpenCensus]:https://opencensus.io
[Fn medium]:https://medium.com/fnproject
