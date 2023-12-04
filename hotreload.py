import sys
import os
import subprocess
import argparse
import time
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler

class MyHandler(FileSystemEventHandler):
    def __init__(self, script_path, script_args, debounce_time=0.1):
        self.script_path = script_path
        self.script_args = script_args
        self.debounce_time = debounce_time
        self.last_event_time = 0
        super().__init__()

        self._rerun()

    def on_modified(self, event):
        if event.src_path.endswith(".py") and (time.time() - self.last_event_time) > self.debounce_time:
            file_path = os.path.basename(event.src_path)
            print(f"\n{file_path} has changed. Reloading...")
            self._rerun()

    def _rerun(self) -> None:
            try:
                cmd = [sys.executable, self.script_path] + self.script_args
                subprocess.run(cmd)
            except Exception as e:
                print(f"Error reloading file: {e}")
            finally:
                self.last_event_time = time.time()

if __name__ == "__main__":
    args = argparse.ArgumentParser(description="Python Hot Reloader")

    args.add_argument("-s" , "--script", type=str, required=True, help="Input Script")

    args = args.parse_args()

    script = args.script

    input_file_path = script.split(" ")[0]
    input_dir_path = os.path.dirname(input_file_path)
    input_script_args = script.split(" ")[1:]

    print("Watching...\n\n")

    event_handler = MyHandler(input_file_path, input_script_args)

    observer = Observer()
    observer.schedule(event_handler, input_dir_path, recursive=False)
    observer.start()

    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        observer.stop()

    observer.join()

