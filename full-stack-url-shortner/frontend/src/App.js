import React, { useState } from "react";
import axios from "axios";

function App() {
  const [longURL, setLongURL] = useState("");
  const [shortURL, setShortURL] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post("http://localhost:8080/shorten", {
        long_url: longURL,
      });
      setShortURL(response.data.short_url);
    } catch (error) {
      console.error("Error shortening URL:", error);
    }
  };

  return (
    <div style={{ padding: "20px" }}>
      <h1>URL Shortener</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Enter long URL"
          value={longURL}
          onChange={(e) => setLongURL(e.target.value)}
          style={{ width: "300px", marginRight: "10px" }}
        />
        <button type="submit">Shorten</button>
      </form>
      {shortURL && (
        <div>
          <h2>Short URL:</h2>
          <a href={shortURL} target="_blank" rel="noopener noreferrer">
            {shortURL}
          </a>
        </div>
      )}
    </div>
  );
}

export default App;
