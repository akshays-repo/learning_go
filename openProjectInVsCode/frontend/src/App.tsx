import { useEffect, useState } from "react";
import "./App.css";
import { GetDirectorys, OpenInVsCode } from "../wailsjs/go/main/App";

type FileCardType = {
  name: string;
  onClick: () => void;
};

const FileCard = ({ name, onClick }: FileCardType) => {
  return (
    <div className="w-96 h-88 p-6 max-w-sm bg-white rounded-lg border border-gray-200 shadow-md dark:bg-gray-800 dark:border-gray-700">
      <a href="#">
        <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
          {name}
        </h5>
      </a>
      <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">
        {/* Here are the biggest enterprise technology acquisitions of 2021 so far,
        in reverse chronological order. */}
      </p>
      <button
        onClick={() => {
          onClick();
        }}
        className="inline-flex items-center py-2 px-3 text-sm font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
      >
        OpenVscode{" "}
        <svg
          aria-hidden="true"
          className="ml-2 -mr-1 w-4 h-4"
          fill="currentColor"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            fill-rule="evenodd"
            d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z"
            clip-rule="evenodd"
          ></path>
        </svg>
      </button>
    </div>
  );
};

const openInVsCode = (dir: string) => {
  OpenInVsCode(dir);
};

function App() {
  const [listDir, setlistDir] = useState<string[]>([]);

  useEffect(() => {
    getAppDirctorys();
  }, []);

  const getAppDirctorys = () => {
    GetDirectorys().then((dir) => {
      console.log(dir);
      setlistDir(dir);
    });
  };

  return (
    <>
      <div id="App bg-gray-50">
        <div className="flex flex-wrap item-center ">
          {listDir.map((dir) => {
            return (
              <div className="m-5">
                {" "}
                <FileCard
                  name={dir}
                  onClick={() => {
                    openInVsCode(dir);
                  }}
                />
              </div>
            );
          })}
        </div>
      </div>
    </>
  );
}

export default App;
