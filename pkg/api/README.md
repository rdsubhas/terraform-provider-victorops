[Victorops swagger schema source](https://portal.victorops.com/public/api-docs.html). Some changes needed to make it compile:

* `#/definitions/OnCallInterval`: Search for keys called `on:` and `off:` and replace with `"on":` and `"off:"`. on/off has a special meaning in yaml. These keys are being parsed as bools instead of strings.

* `#/parameters/noteName`: Add `in: path`

* `#/paths/"api-public/v1/incidents"/post`: Typo, change `#/definitions/Note` to `#/definitions/Notes`

* in all parameters, `required: true|false` should not be present. Find and comment it out.
