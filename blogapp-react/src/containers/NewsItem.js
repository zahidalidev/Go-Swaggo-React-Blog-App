import React from "react";
import { useSelector } from "react-redux";

const imgStyle = {
  hight: "auto",
  width: "80%",
  border: "4px solid RebeccaPurple ",
  borderRadius: "5%",
};

const articleStyle = {
  width: "50%",
  margin: "0 auto",
  color: "olive",
};

let NewsItem = () => {
  const article = useSelector((state) => state.news);
  return article ? (
    <>
      {article.map((item) => (
        <article style={articleStyle}>
          <div>
            <h1>{item.title}</h1>
            <img style={imgStyle} src={item.urlToImage} alt="" />
            <h4>{item.description}</h4>
            <a href={item.url} target="_blank">
              READ MORE
            </a>
          </div>
        </article>
      ))}
    </>
  ) : null;
};

export default NewsItem;
