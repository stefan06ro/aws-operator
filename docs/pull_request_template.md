<!--
Every PR must be tested with `awscnf` before merging.

in oder to execute `awscnf` with your changes do folowing:
* download or build latest `awscnf`
* choose a test installation (ie. ginger)
* deploy your custom release with: `opsctl release deploy -i ginger -b 1x.x.x -c aws-operator@YOUR_BRANCH`
* create kubeconfig to your test installation `opsctl create kubeconfig -i ginger`
* set release version for `awscnf`  with `export AWSCNFM_CREATE_RELEASEVERSION=YOUR_RELEASE`
* execute `awscnf plan pl001`

When `awscnf` succeed, post the result logs in the PR.
-->

## Checklist

- [ ] Update changelog in CHANGELOG.md.
- [ ] I have sucesfully executed `awscnf plan pl001` with custom release with aws-operator
