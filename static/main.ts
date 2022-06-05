import { Fzf } from "./node_modules/fzf/dist/fzf.es.js";

main();

async function fetchPath(url: string) {
  return await fetch(url)
    .then((response) => {
      return response.json();
    })
    .catch((response) => {
      return Promise.reject(
        new Error(`{${response.status}: ${response.statusText}`),
      );
    });
}

function fzfSearch(list: string[], keyword: string): string[] {
  const fzf = new Fzf(list);
  const entries = fzf.find(keyword);
  const ranking: string[] = entries.map((entry: Fzf) => entry.item);
  return ranking;
}

async function main() {
  const url: URL = new URL(window.location.href);
  // サーバーからExcelの行を取得
  const list: Promise<string[]> = await fetchPath(url.origin + "/list");
  const searchInput: HTMLElement = document.getElementById("search-form");
  const resultOutput: HTMLElement = document.getElementById("search-result");
  // キーを押すたびにページ内容更新
  searchInput.addEventListener("keyup", () => {
    // 要素クリア
    while (resultOutput.firstChild) {
      resultOutput.removeChild(resultOutput.firstChild);
    }
    // fzf検索
    const result = fzfSearch(list, searchInput.value);
    // 検索結果をコンソールに表示
    console.log(result);
    // 検索結果を結果要素に表示
    result.map((line: string) => {
      const p = document.createElement("p");
      const text = document.createTextNode(line);
      p.appendChild(text);
      resultOutput.append(p);
    });
  });
}
