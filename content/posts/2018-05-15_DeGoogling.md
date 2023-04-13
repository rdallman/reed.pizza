---
title: DeGoogling
description: In depth guide for freeing myself from the chains of Google
date: '2018-05-15T07:50:41.980Z'
categories: []
keywords: []
---

![](/1__0cF__d1AMwh3X__eACJXzhMw.gif#center)

I have been rdallman10@gmail.com since AIM wasn’t cool anymore, for over 10 years now — my whole ‘adult’ life. When the HTC Evo 4g came out, that same day I abandoned my Blackberry and started installing random kernels from xda-developers.com in an attempt to turn my phone into an object that could get hot enough to allow me to continue the work of the late, great Dr. Octavius. It was that glorious, fateful day I fell into the Google suck zone.

![](/1__UFxiSq8Wy8YRpYbB6oY31Q.gif#center)

For getting through college, having a gmail account was undoubtedly awesome; Docs, Drive, Calendar, Gmail, Maps, and Search are still things that I think Google is doing better than anyone else, and for free! Early days Maps was a wild ride, and I often led family and friends to destinations that did not seem to exist. I will never live some of those stories down, and in the more interesting ones, I’m lucky to have lived! Now, thanks to Apple Maps, I am once again some kind of derelict Magellan, and order has been restored to my Constanza-esque galaxy of frequent, minor inconveniences.

As of 5/14/18 my google account ceases to exist, yet life continues. I’d like to illustrate exactly what it took. I found these kinds of rare posts useful when I was myself migrating off of Google, so I’d like to make another one to help out other people contributing to the shortage of tin foil that is sweeping the nation.

#### Tid bits:

*   It took me 36 days. I managed to keep my job, for now. The wedding invitations weren’t exactly flowing in, either. Message received, “friends”.
*   Do not do this for your mother/aunt/grandpa or you’ll be on the hook for a lot more than fixing the WiFi over Thanksgiving.
*   I include no affiliate links to avoid giving off the impression that I have capitalist motives, like most people that blog about services (you heard it here first…). On a related note, doing all of this cost me north of $1500, some of that recurring annually; $obligatory\_2018\_privilege\_disclaimer
*   I would be just delighted to hear any criticisms or alternatives if anybody manages to read this, ya don’t know what ya don’t know.
*   If you enjoy your current stream of consciousness, abandon hope all ye who enter here.

### The Great Byte Migration

The EFF has a pretty fantastic web page [https://ssd.eff.org/en](https://ssd.eff.org/en) that goes over a lot of the stuff here in greater detail, with tutorials. This post is mostly meant to document the insane amount of work required to replace Google along with my choices when doing so to aid others, it is not meant to exhaustively detail the horrors of allowing Facebook and Google to harvest human data under the guise of serving them. My threat model is likely not sufficient for people that are out there actually attempting to make the world a better place, like journalists or others with sensitive information.

Mostly, I am done tacitly consenting to training robots and at this point I know better and can’t live with myself continuing to do so. I am attempting to spread my data across multiple providers at a minimum, own my data where possible, necessary and/or reasonable and hide my communications simply on principle, since I believe we should have privacy when talking to others. Hopefully, one day I can actually do my part in building things that make this easier for lay people (if I gleaned anything valuable from this exercise, it’s that). This is a big step for me in this dept., but this is likely a lifelong journey and I don’t feel like it’s complete, even now. War is peace.

#### TL;DR

*   gmail -> FastMail + custom domain
*   Pixel 2 -> iPhone X
*   Project Fi -> T-Mobile
*   8.8.8.8 -> 1.1.1.1
*   ExpressVPN -> NordVPN
*   Dropbox -> Tresorit
*   Google search -> DuckDuckGo
*   Hangouts -> Signal, iMessage
*   rm -rf facebook instagram

#### Email:

*   Custom domain (**gandi.net**: 2FA, no Super Bowl ads to fund evil), with **Cloudflare** for easy DNSSEC, SPF, DKIM (Acronyms!) + easy https:// for web page. I am now and forever reed@rdallman.com and I can carry this around to other email providers. Not that @gmail is dying anytime soon.
*   **Fastmail** for email, calendar, contacts. Email itself in its current form is not easily secured, I do not intend to make a point of this here (they are in a 5 eyes country, which I don’t like, but I pretty much live knowing somebody else can read most of my emails if they really want to, and work around that). One of my constraints was that this would work with the native iOS email, calendar **and** contacts, for which Protonmail is not an option. Protonmail is cool, but for my every day email it’s not there at least for now. Posteo was the other option I liked, and they’re a really cool company, but no custom domain, and I don’t ever want to do this again. Fastmail is doing lots of work on making email, calendars, and open standards around those things better, which is something I like paying money for. ❤
*   Gmail does not allow recycling of email addresses, so I don’t have to worry about nefarious little snots becoming rdallman10@gmail.com and trying to tell my bank about it (at least, until Google goes under). Migrating my email, contacts, and calendar out was a breeze (mostly, thanks to Fastmail). Changing my email across all my accounts was painful, according to 1Password I had 250 accounts at the start (down to 160 now) — I’m glad I’ll never have to do this again, some companies don’t even allow this and I had to go through their support on some.
*   Snooze + reminders in Inbox was life. I have IMAP folders now. It’s more of a fucking lie than red velvet. Hopefully, this feature comes to all email. I looked at Newton for email, but it is pretty privacy invasive to get my beloved snooze back, and isn’t native to iOS (I kinda hate 3rd party apps for email, since they die every few years).

#### Phone:

*   I bought an **iPhone X**. I was as reluctant as the next pimple ridden nerd, I’ve been on Android since the HTC Evo 4g in 2009 without an iPhone blip and now at last Android finally doesn’t suck the last few years. It’s kinda like a fraternity getting kicked off of campus the semester after you get through all the hazing — all that suffering, no payout. The iPhone experience has mostly been really nice so far, aside from notifications completely sucking in comparison. I think I still prefer Android, but this is worth the tradeoffs, and the iPhone does look and feel a lot better in my hand than my Pixel 2XL when comparing them. They also seem capable of sticking to one messaging app.
*   Apple seems to be taking privacy seriously, which may be the 2nd biggest honeypot of all time (thanks Google), but I’ll roll with it. Getting my family and friends all on Signal would be harder than getting them all to agree on if the moon is made of cheese (some of us haven’t lost hope), so iMessage is pretty convenient for securely chatting with them. Still don’t approve of using biometrics as secrets, and training the robots with our faces we are so doomed — I’ll use a pin, but it’s not going to save us (nor is it perfect, thanks CCTV). I also disable most of iCloud and I only use it for Apple Pay, which I find confusing as a requirement, but I’ll oblige. Android, on the other hand, is completely broken for privacy, with the Play Store tracking your location for whatever reason. I looked at using Android without Play Store, no thank you, aside from other shortcomings. Switching to iPhone was likely a bigger privacy win than switching my email, even. The amount of data I downloaded from Google was disgusting, and most of that is because I used Android for years.
*   iOS calendar/contacts/email works great with Fastmail, and I can add an event from an email with the push of a button, without needing a gmail account. I use this a lot. By comparison, Android vanilla experience without using gcal/gmail/gcontacts totally sucks, you have to use CardDav sync and CalDav sync apps, make sure your contacts go to the right set — I’ve had more fun clipping ingrown toe nails — that’s not to mention K9 and alternatives. No, thanks. This was a huge selling point for iPhone for me.
*   iOS I can use **DuckDuckGo** as the default search engine. On Android, I had to take extra care to click the DDG box instead of the Google box to search, it’s nice to have an integrated experience.
*   Not only did Google have my phone, they also had my phone service. While Project Fi was exceptionally awesome, especially for international travel, and perfect for me since I’m mostly on WiFi for data anyway, I am amazed at how bad Google had me by having both my phone and phone service. I’ve switched to **T-Mobile** which offers similar service with international data (but throttled, boo) for about double the price. I’m pretty bummed about this one, and have no illusions about T-Mobile having an iota of a care for privacy, either (or any TelCo, unfortunately). Google had all my voicemails, call records and texts for the last 8 years, apparently T-Mobile only retains this for 1 year (so they say), so in some regards they’re a little better. Side note, my texts from 8 years ago are bad and I am so sorry for being me and thanks for even conversing with me and I’m working on it.
*   RIP Cast. I was using my phone as my remote for some time now, and it was really great vs. the old school tv remote (search is life changing). Some iOS apps seem to have the ability to cast to my Google TV, but I am on the fence about getting rid of my Chromecast and Google TV still — at least, now I’m logged out, but I’m still wary. Google Assistant by now knows I only use it to set a mac n cheese timer, anyway, but there’s no telling what else she might know.

**Drive/Dropbox:**

*   Dropbox also got me through college and I used it pretty heavily for backing up my stuff, keeping a list of my random, awful sitcom ideas, and storing my photos. Their product was really the first of its kind, but now there are secure alternatives with the same features but that don’t have a way for their employees or gov. employees to be able to access your data. I’m now using **Tresorit**, which is so far so good, they have a file upload size limit which kind of sucks, so I had to unzip some zips to get them in, but they have all the features I used from dropbox: camera upload, sync, offline files, in app editor, URL sharing — but they allege to not have access to your data using end-to-end encryption, I’ll bite. This one was the easiest thing to migrate and honestly there is no excuse to continue to use Dropbox, even at work it’s time to convince people to use secure alternatives.
*   For documents, I’m using Pages, Numbers & Keynote on my Macbook now, backed up to my Tresorit, which is fine. I still like Google Docs/Sheets/Slides a lot, it was my default, so this is taking some adjusting to (one preso down!).

#### VPN:

*   I was using ExpressVPN, which only offered OpenVPN configurations. OpenVPN is something I trust a little more than IKEv2, but IKEv2 is built in natively on iOS and macOS, and allow me to use the “Connect on Demand” feature to get an always on VPN. For me, this is something I was doing on Android for years with OpenVPN, so I went looking for an IKEv2 friendly VPN. **NordVPN** offers configuration of IKEv2 in the native settings for iOS and macOS, and this has been going smoothly. It’s cheaper and I haven’t had any speed issues thus far, but I’ve used their service before and had gripes with the speed, especially when traveling. We’ll see how this goes. Using a VPN all the time can be painful, some things block it (Netflix, Amazon), and Android had a nice way of adding whitelist for certain apps which I haven’t found on iOS, but this is something I’m committed to. If you’ve never done this, it’s really not hard to get set up and is relatively cheap with a decent return on investment.

#### Search:

*   **DuckDuckGo**, as mentioned earlier, is a pretty easy win to migrate to, and integrates well into Firefox on desktop and as the default search on iOS. I’ve learned to bang things, thanks DDG. Some things like movie showtimes I will usually !g, Google is kinda light years ahead of the competition on some stuff and I’ve been spoiled, but I try to stick to the DDG searches as much as I can.

#### DNS:

*   You already know my DNS was set to 8.8.8.8 — thanks to recent Cloudflare **1.1.1.1** I’ve changed this, too. It’s free, so surely they’re up to something, but at least it’s not Google I guess. They’re also working on DNS over TLS, which I’m in for.

#### Hangouts:

*   Google can’t stick with a messaging app anyway, but I’ve used Hangouts for over 4 years now and of course still use it to chat with some people. Hopefully, one day we’ll all agree on something relatively secure, but at the same time, competition is likely good to avoid us having one giant honeypot. I am pushing my friends to use **Signal**, but this is an uphill battle. It’s the easiest thing to install and use, and the UI is really great on both Android and iOS, but while people seem to have no issue loading 40 versions of Bejeweled into their phone, one more chat app to talk to 1 person seems like too much to ask. The war wages on. For many people, I will gladly use iMessage now, it’s vastly superior to SMS which I was using for many of them previously. Interestingly, this is mostly an issue for my tech friends who are on Android :)

#### Facebook/Instagram:

![](/1__guGFeN7I__PHsU__znHaCIHw.gif#center)

*   Also took care of this. This one was easy, too. I haven’t used it very much since high school when it was the de facto messaging app. I am so glad to erase those messages from the ether. I am still in search of a new platform to share my beach photos on to seek my ever needed validation and to discourage adolescent obesity. I expect even fewer wedding invites now, as well. For as much of a web presence as I work to maintain, for non-tech people Facebook really is the only thing many people use to locate people, and people use it for events too. However, I think that usage will phase out, younger generations are not using Facebook at all, so long term I don’t feel like I’m missing very much here.

### The End

This sucked, was probably not worth it in the greater cosmic picture (what is?) especially since I surely convinced 0 people to follow suit, but was worth it to me for peace of mind (though, it didn’t buy much). At least, it was interesting to analyze all the ways Google completely owns me, and after doing that I felt more motivated to get out. Part of this exercise was definitely that I know better and am capable of making this change, and I hope that the engineers building all our software will consider making this all easier for lay people, myself included, and baking in privacy by default to all their software. I feel somewhat morally bankrupt as it is for not devoting my time to working on making any of this software better, since I’m capable of doing so and aware of the issue(s). There are still some things I’d like to fix personally, such as getting off of Amazon (Kindle is my weakness), and credit cards are pretty much the next worst thing for privacy — they’re selling all of our purchases to anybody willing to pay for them — but this was a decent step for me. Unfortunately, I still expect the robots to enslave us humans in my lifetime and I’ve either moved myself into the first wave for my now public mutiny or I’ve simply saved myself for the second wave. Ignorance is strength, comrades.
