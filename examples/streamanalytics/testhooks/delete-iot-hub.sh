#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Delete the StreamInputIOTHub resource before deleting secret with password
${KUBECTL} delete streaminputiothub.streamanalytics.azure.upbound.io --all