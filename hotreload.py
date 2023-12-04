"""
Hot Reload for a Python Script.
"""

import sys
import os
import subprocess
import argparse
import time
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler

class ReloadHandler(FileSystemEventHandler):
    """
    Reload Handler for the script. 

    on_modified is called when the script is modified.
    """

    def __init__(self, script_path: str, script_args: list[str], debounce_time: float = 0.1):
        """
        constructor.

        script_path: path to the script to be reloaded.
        script_args: arguments to be passed to the script.
        debounce_time: minimum time between reloads.
        """

        self._script_path: str = script_path
        self._script_args: list[str] = script_args

        # to prevent multiple reloads in a short time
        self._debounce_time: float = debounce_time
        self._last_event_time: float = 0

        super().__init__()

        self._reload()

    def on_modified(self, event) -> None:
        """
        Called when the script is modified.

        event: event object.
        """

        if not event.src_path.endswith(".py") or self._is_too_soon():
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
            cmd: list[str] = [sys.executable, self._script_path] + self._script_args
            subprocess.run(cmd)
        except Exception as e:
            print(f"Error reloading file: {e}")
        finally:
            self._last_event_time = time.time()


if __name__ == "__main__":
    args = argparse.ArgumentParser(description="Python Hot Reloader")

    args.add_argument("-s" , "--script", type=str, required=True, help="Input Script")

    args = args.parse_args()

    script = args.script

    input_file_path = script.split(" ")[0]
    input_dir_path = os.path.dirname(input_file_path)
    input_script_args = script.split(" ")[1:]

    print("Watching...\n\n")

    # create event handler
    event_handler = ReloadHandler(input_file_path, input_script_args)

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

