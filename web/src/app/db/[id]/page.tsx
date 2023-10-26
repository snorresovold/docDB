import axios from "axios";

export interface Collection {
  name: string;
  documents: Document[];
}

export interface Document {
  id: string;
  content: any;
}

async function getCollection(collection: string): Promise<Collection> {
  try {
    const response = await axios.get<Collection>(
      `http://localhost:8080/getCollection/${collection}`
    );
    console.log(response.data);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export default async function CollectionPage({ params }: any) {
  try {
    const documents = await getCollection(params.id);

    return (
      <div className="p-4">
        <h1 className="text-2xl font-bold mb-4">{documents.name}</h1>
        {documents.documents?.map((e) => {
          return (
            <div key={e.id} className="mb-4 p-4 border border-gray-200">
              <h2 className="text-xl font-bold mb-2">{e.id}</h2>
              <p>{JSON.stringify(e.content)}</p>
            </div>
          );
        })}
      </div>
    );
  } catch (error) {
    // Handle the error in the UI or log it
    console.error(error);
    return <div className="p-4">Error fetching collection</div>;
  }
}
