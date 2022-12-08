def pytest_report_teststatus(report):
    if report.when == "call":
        print(f"duration: {int(report.duration * 1000)}ms")
