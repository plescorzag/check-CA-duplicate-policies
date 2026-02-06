This README provides clear instructions on how to extract your Kubernetes CA bundle and use the `check_bundle.go` script to identify which certificate is causing the `x509` validation error.

---

# Certificate Bundle Validator
This utility helps troubleshoot the **"x509: invalid certificate policies"** or **"duplicate OIDs"** errors often encountered in Go-based Kubernetes Operators. It parses a PEM bundle and validates each certificate individually using Go's standard `crypto/x509` library.

## 1. Prerequisites
* **Go installed** (v1.19+ recommended for strict validation checks).
* **oc** configured to access your OpenShift cluster.
* **OpenSSL** (optional, for manual inspection).

## 2. Setup
Download the `check_bundle.go`code

## 3. How to get the CA Bundle
To extract the root CA bundle from your OpenShift cluster, run:

```bash
# Extracts the CA from the default ConfigMap in your current namespace
oc get configmap kube-root-ca.crt -o jsonpath='{.data.ca\.crt}' > ca-all.crt
```

## 4. Usage
Run the validator against the extracted file:
```bash
go run check_bundle.go ca-all.crt
```

### Expected Output
The script will list every certificate found in the file. If a certificate violates **RFC 5280** (like having duplicate Policy OIDs), it will mark that specific entry as **FAIL** and print the exact reason.

---

## 5. Possible Output
| Error Message | Meaning |
| --- | --- |
| `x509: certificate contains duplicate extensions` | An OID (like Certificate Policies) appears more than once. This is forbidden in Go 1.19+. |
| `x509: invalid certificate policies` | The structure inside the Policy extension is malformed (e.g., empty qualifiers). |
| `File Error: open ca-all.crt: no such file...` | The file wasn't created correctly or you are in the wrong directory. |

