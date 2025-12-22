"""
Hot Reload for a Script (Python, Ruby, etc.).
"""

import argparse
import os
import subprocess
import sys
import time

from watchdog.events import FileSystemEventHandler
from watchdog.observers import Observer


class ReloadHandler(FileSystemEventHandler):
    """
    Reload Handler for the script.

    on_modified is called when the script is modified.
    """

    def __init__(
        self,
        script_path: str,
        script_args: list[str],
        runner: str | None = None,
        debounce_time: float = 0.1,
    ):
        """
        constructor.

        script_path: path to the script to be reloaded.
        script_args: arguments to be passed to the script.
        runner: command used to run the script (e.g. 'python', 'ruby').
        debounce_time: minimum time between reloads.
        """

        self._script_path: str = script_path
        self._script_args: list[str] = script_args

        # if runner is not provided, infer from extension or default to current python
        if runner is not None:
            self._runner = runner
        else:
            if script_path.endswith(".rb"):
                self._runner = "ruby"
            elif script_path.endswith(".py"):
                self._runner = sys.executable
            else:
                self._runner = sys.executable

        # to prevent multiple reloads in a short time
        self._debounce_time: float = debounce_time
        self._last_event_time: float = 0

        super().__init__()

        self._reload()

    def on_modified(self, event) -> None:
        """
        Called when a file is modified.

        event: event object.
        """

        # only react to changes to the target script itself
        if os.path.abspath(event.src_path) != os.path.abspath(self._script_path):
            return

        if self._is_too_soon():
            return

        file_path = os.path.basename(event.src_path)

        print(f"\n{file_path} has changed. Reloading...")

        self._reload()

    def _is_too_soon(self) -> bool:
        """
        Checks if the last reload was too soon.
        """
        return (time.time() - self._last_event_time) <= self._debounce_time

    def _reload(self) -> None:
        """
        Reloads the script by running a new subprocess.
        """

        try:
            cmd: list[str] = [self._runner, self._script_path] + self._script_args
            subprocess.run(cmd)
        except Exception as e:
            print(f"Error reloading file: {e}")
        finally:
            self._last_event_time = time.time()


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Script Hot Reloader")

    parser.add_argument(
        "-s",
        "--script",
        type=str,
        required=True,
        help='Script command, e.g. "path/to/main.rb arg1 arg2"',
    )

    parser.add_argument(
        "-r",
        "--runner",
        type=str,
        default=None,
        help="Command to run the script (e.g. 'python', 'python3', 'ruby'). "
        "If omitted, it is inferred from the script extension.",
    )

    args = parser.parse_args()

    script = args.script

    input_file_path = script.split(" ")[0]
    input_dir_path = os.path.dirname(input_file_path) or "."
    input_script_args = script.split(" ")[1:]

    print(f"Watching {input_file_path}...\n")

    # create event handler
    event_handler = ReloadHandler(
        input_file_path,
        input_script_args,
        runner=args.runner,
    )

    # start observer for the script's directory
    observer = Observer()

    # use the event handler to re-run on file changes in the directory
    observer.schedule(event_handler, input_dir_path, recursive=False)
    observer.start()

    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        observer.stop()

    observer.join()
