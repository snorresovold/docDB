import React from "react";

interface Collection {
  name: string;
}

function Collection({ name }: Collection) {
  return (
    <div>
      <p>{name}</p>
    </div>
  );
}

export default Collection;
