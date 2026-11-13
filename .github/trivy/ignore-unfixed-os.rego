package trivy

import rego.v1

default ignore := false

# Ignore Debian OS-package CVEs that have no fix available yet, up to severity HIGH.
# CRITICALs are never ignored. As soon as Debian ships a fix, the status changes to
# "fixed" and the finding is reported again automatically.
#
# Used by .github/workflows/reusable-cve-scan.yml (input: ignore_unfixed_os).
# Keep in sync with the stale-entry check in that workflow.
ignore if {
	startswith(input.PkgIdentifier.PURL, "pkg:deb/")
	input.Status in {"affected", "fix_deferred", "will_not_fix", "end_of_life"}
	input.Severity != "CRITICAL"
}
