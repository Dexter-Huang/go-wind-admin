"""检查所有 tx7do 模块均解析到仓库源码 / Check local module replacements."""

import json
from pathlib import Path
import subprocess


def main():
    backend = Path(__file__).resolve().parents[1]
    source = backend / "third_party" / "tx7do"
    entries = json.loads((source / "modules.json").read_text(encoding="utf-8"))
    expected = {entry["module"]: entry for entry in entries}
    if len(expected) != len(entries):
        raise SystemExit("Duplicate modules in modules.json")
    output = subprocess.check_output(
        ["go", "list", "-mod=readonly", "-m", "-json", "all"],
        cwd=backend, text=True, encoding="utf-8",
    )
    decoder = json.JSONDecoder()
    found = set()
    while output.strip():
        module, end = decoder.raw_decode(output.lstrip())
        output = output.lstrip()[end:]
        name = module["Path"]
        if not name.startswith("github.com/tx7do/"):
            continue
        if name not in expected:
            raise SystemExit(f"Module has not been imported: {name}")
        entry = expected[name]
        directory = (source / entry["directory"]).resolve()
        if not directory.is_relative_to(source.resolve()):
            raise SystemExit(f"Directory is outside source tree: {directory}")
        replacement = module.get("Replace", {})
        if (not replacement.get("Dir") or replacement.get("Version")
                or Path(replacement["Dir"]).resolve() != directory):
            raise SystemExit(f"Module does not use repository source: {name}")
        if module.get("Version") != entry["version"]:
            raise SystemExit(f"Version differs from provenance manifest: {name}")
        if not (directory / "go.mod").is_file() or not list(directory.glob("LICENSE*")):
            raise SystemExit(f"Missing go.mod or LICENSE: {name}")
        found.add(name)
    if found != expected.keys():
        raise SystemExit(f"Manifest modules missing from graph: {sorted(expected.keys() - found)}")
    print(f"OK: all {len(found)} tx7do modules resolve to repository source.")


if __name__ == "__main__":
    main()
