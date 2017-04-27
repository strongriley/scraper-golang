# scraper-golang
Rewritten scraper in golang

## Using

```
go build
./scraper-golang -url http://example.com -n [WORKERS] -m [MAX_URLS]
```

## Example Output

```
[
{
  "url": "http://github.com/strongriley",
  "static_urls": [
    "https://assets-cdn.github.com/assets/frameworks-d7b19415c108234b91acac0d0c02091c860993c13687a757ee345cc1ecd3a9d1.css",
    "https://assets-cdn.github.com/assets/github-97f8afbdb0a810d4ffa14a5fc7244b862b379d2c341d5eeb89489fbd310e4a83.css",
    "https://assets-cdn.github.com/assets/site-537c466d44a69d38c4bd60c2fd2955373ef96d051bd97b2ad30ed039acc97bff.css",
    "https://avatars0.githubusercontent.com/u/313228?v=3\u0026s=64",
    "https://avatars0.githubusercontent.com/u/313228?v=3\u0026s=460",
    "https://assets-cdn.github.com/images/spinners/octocat-spinner-32.gif",
    "https://assets-cdn.github.com/images/spinners/octocat-spinner-128.gif",
    "https://assets-cdn.github.com/assets/compat-8a4318ffea09a0cdb8214b76cf2926b9f6a0ced318a317bed419db19214c690d.js",
    "https://assets-cdn.github.com/assets/frameworks-6d109e75ad8471ba415082726c00c35fb929ceab975082492835f11eca8c07d9.js",
    "https://assets-cdn.github.com/assets/github-a6e170e1a1432572f27d4c18936acbd70b123dc8cb36302095396fb6b2096b32.js"
  ],
  "linked_urls": [
    "http://github.com/strongriley",
    "https://github.com",
    "http://github.com/features",
    "http://github.com/business",
    "http://github.com/explore",
    "http://github.com/pricing",
    "http://github.com/login?return_to=%2Fstrongriley",
    "http://github.com/join?source=header",
    "http://github.com/contact/report-abuse?report=strongriley",
    "http://github.com/strongriley?tab=repositories",
    "http://github.com/strongriley?tab=stars",
    "http://github.com/strongriley?tab=followers",
    "http://github.com/strongriley?tab=following",
    "http://github.com/strongriley/riddler",
    "http://github.com/strongriley/wedding",
    "http://github.com/strongriley/d3",
    "http://github.com/d3/d3",
    "http://github.com/strongriley/d3/stargazers",
    "http://github.com/strongriley/d3/network",
    "http://github.com/strongriley/python-dbase",
    "http://github.com/strongriley/python-dbase/stargazers",
    "http://github.com/strongriley/python-dbase/network",
    "http://github.com/strongriley/remove-tumblr-redirects",
    "http://github.com/strongriley?tab=overview\u0026from=2017-04-01\u0026to=2017-04-27",
    "http://github.com/strongriley?tab=overview\u0026from=2016-12-01\u0026to=2016-12-31",
    "http://github.com/strongriley?tab=overview\u0026from=2015-12-01\u0026to=2015-12-31",
    "http://github.com/strongriley?tab=overview\u0026from=2014-12-01\u0026to=2014-12-31",
    "http://github.com/strongriley?tab=overview\u0026from=2013-12-01\u0026to=2013-12-31",
    "http://github.com/strongriley?tab=overview\u0026from=2012-12-01\u0026to=2012-12-31",
    "http://github.com/strongriley?tab=overview\u0026from=2011-12-01\u0026to=2011-12-31",
    "http://github.com/strongriley?tab=overview\u0026from=2010-12-01\u0026to=2010-12-31",
    "http://github.com/strongriley?tab=overview\u0026from=2010-06-01\u0026to=2010-06-30",
    "http://github.com/strongriley/housing",
    "http://github.com/strongriley/housing/commits?author=strongriley\u0026since=2017-04-01T00:00:00Z\u0026until=2017-04-28T00:00:00Z",
    "http://github.com/strongriley/scraper-golang",
    "http://github.com/strongriley/scraper-golang/commits?author=strongriley\u0026since=2017-04-01T00:00:00Z\u0026until=2017-04-28T00:00:00Z",
    "https://github.com/contact",
    "https://github.com/blog",
    "https://github.com/about",
    "https://github.com/site/terms",
    "https://github.com/site/privacy",
    "https://github.com/security"
  ]
},
{
  "url": "https://github.com",
  "static_urls": [
    "https://assets-cdn.github.com/assets/frameworks-d7b19415c108234b91acac0d0c02091c860993c13687a757ee345cc1ecd3a9d1.css",
    "https://assets-cdn.github.com/assets/github-97f8afbdb0a810d4ffa14a5fc7244b862b379d2c341d5eeb89489fbd310e4a83.css",
    "https://assets-cdn.github.com/assets/site-537c466d44a69d38c4bd60c2fd2955373ef96d051bd97b2ad30ed039acc97bff.css",
    "https://assets-cdn.github.com/images/modules/site/satellite-logo.png",
    "https://assets-cdn.github.com/images/modules/site/satellite-wordmark.png",
    "https://assets-cdn.github.com/images/modules/site/home-illo-conversation.svg",
    "https://assets-cdn.github.com/images/modules/site/home-illo-chaos.svg",
    "https://assets-cdn.github.com/images/modules/site/home-illo-business.svg",
    "https://assets-cdn.github.com/images/modules/site/integrators/slackhq.png",
    "https://assets-cdn.github.com/images/modules/site/integrators/zenhubio.png",
    "https://assets-cdn.github.com/images/modules/site/integrators/travis-ci.png",
    "https://assets-cdn.github.com/images/modules/site/integrators/atom.png",
    "https://assets-cdn.github.com/images/modules/site/integrators/circleci.png",
    "https://assets-cdn.github.com/images/modules/site/integrators/codeship.png",
    "https://assets-cdn.github.com/images/modules/site/integrators/codeclimate.png",
    "https://assets-cdn.github.com/images/modules/site/integrators/gitterhq.png",
    "https://assets-cdn.github.com/images/modules/site/integrators/waffleio.png",
    "https://assets-cdn.github.com/images/modules/site/integrators/heroku.png",
    "https://assets-cdn.github.com/images/modules/site/logos/airbnb-logo.png",
    "https://assets-cdn.github.com/images/modules/site/logos/sap-logo.png",
    "https://assets-cdn.github.com/images/modules/site/logos/ibm-logo.png",
    "https://assets-cdn.github.com/images/modules/site/logos/google-logo.png",
    "https://assets-cdn.github.com/images/modules/site/logos/paypal-logo.png",
    "https://assets-cdn.github.com/images/modules/site/logos/bloomberg-logo.png",
    "https://assets-cdn.github.com/images/modules/site/logos/spotify-logo.png",
    "https://assets-cdn.github.com/images/modules/site/logos/swift-logo.png",
    "https://assets-cdn.github.com/images/modules/site/logos/facebook-logo.png",
    "https://assets-cdn.github.com/images/modules/site/logos/node-logo.png",
    "https://assets-cdn.github.com/images/modules/site/logos/nasa-logo.png",
    "https://assets-cdn.github.com/images/modules/site/logos/walmart-logo.png",
    "https://assets-cdn.github.com/assets/compat-8a4318ffea09a0cdb8214b76cf2926b9f6a0ced318a317bed419db19214c690d.js",
    "https://assets-cdn.github.com/assets/frameworks-6d109e75ad8471ba415082726c00c35fb929ceab975082492835f11eca8c07d9.js",
    "https://assets-cdn.github.com/assets/github-a6e170e1a1432572f27d4c18936acbd70b123dc8cb36302095396fb6b2096b32.js"
  ],
  "linked_urls": [
    "https://github.com",
    "https://github.com/features",
    "https://github.com/business",
    "https://github.com/explore",
    "https://github.com/pricing",
    "https://github.com/dashboard",
    "https://github.com/login",
    "https://github.com/join?source=header-home",
    "https://github.com/open-source",
    "https://github.com/join?source=button-home",
    "https://github.com/join?plan=business\u0026setup_organization=true\u0026source=business-page",
    "https://github.com/integrations",
    "https://github.com/personal",
    "https://github.com/join",
    "https://github.com/about",
    "https://github.com/blog",
    "https://github.com/business/customers",
    "https://github.com/about/careers",
    "https://github.com/about/press",
    "https://github.com/contact",
    "https://github.com/site/terms",
    "https://github.com/site/privacy",
    "https://github.com/security"
  ]
}
]
