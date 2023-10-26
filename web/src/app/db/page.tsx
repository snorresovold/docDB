import Collection from "@/components/Collection";
import axios from "axios";
import React from "react";

async function getCollections(): Promise<string[]> {
  try {
    const response = await axios.get<string[]>(
      "http://localhost:8080/getCollections"
    );
    console.log(response.data);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export default async function db() {
  const collections = await getCollections();
  return (
    <div>
      {collections && collections.length > 0 ? (
        collections.map((collection: string) => (
          <Collection name={collection} />
        ))
      ) : (
        <>You have no collections</>
      )}
    </div>
  );
}
