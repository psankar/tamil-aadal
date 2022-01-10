import React, { useState, useEffect } from "react";
import { get as dbget, set as dbset } from "lockr";
import "./App.css";
import Header from "./components/Header";
import Workbench from "./components/Workbench";
import Instructions from "./components/Instructions";

function App() {
  const [hideInstructions, setHideInstructions] = useState(
    dbget("hideInstructions")
  );
  useEffect(() => {
    dbset("hideInstructions", hideInstructions);
  }, [hideInstructions]);

  return (
    <div
      style={{ maxWidth: "600px" }}
      className="is-flex is-flex-direction-column mx-auto"
    >
      <Header onShowInstructions={() => setHideInstructions(false)} />
      {hideInstructions ? (
        <Workbench />
      ) : (
        <Instructions onHide={() => setHideInstructions(true)} />
      )}
    </div>
  );
}

export default App;
