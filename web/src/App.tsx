import "./App.css";

function App() {
  console.log("logging", import.meta.env.VITE_VERSION);
  console.log(import.meta.env);
  return (
    <>
      <div>
        <div>Hello world!</div>
        <div>
          version: {import.meta.env.VITE_VERSION}{" "}
          <pre>{JSON.stringify(import.meta.env)}</pre>
        </div>
      </div>
    </>
  );
}

export default App;
