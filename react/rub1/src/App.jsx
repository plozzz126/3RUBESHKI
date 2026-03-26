import { useReducer, useEffect } from "react";

const initialState = {
    users: [],
    loading: false,
    error: null,
};

function reducer(state, action){
    switch (action.type){
        case "Fetch_Start":
            return {...state, loading: true, error: null};
        case "Fetch_Success":
            return {...state, loading: false, users: action.payload};
        case "Fetch_Error":
            return {...state, loading: false, error: action.payload};
        default :
        return state
    }
}

export default function App(){
  const [state, dispatch] = useReducer(reducer, initialState);
  const {users, loading, error} = state;
  const fetchUsers = async () => {
    dispatch ({type: "Fetch_Start"});
    try {
      const res = await fetch("https://jsonplaceholder.typicode.com/users");
      if (!res.ok) throw new Error(`ошибка: ${res.status}` );
      const data = await res.json();
      dispatch({type: "Fetch_Success", payload: data});  
    } catch (err) {
      dispatch({type: "Fetch_Error", payload: err.message});
    }
  };

  useEffect(() =>{
    fetchUsers();
  }, []);

  return (
    <>
      <div className="da">
        <h1> список пользователей </h1>

        <button onClick={fetchUsers} disabled={loading}>
          {loading ? "загрузка" : "обновить"}
        </button>
        {error && <p>{error}</p>}

        <div className="da2">
          {users.map(user => (
            <div key={user.id} >
              <h3>{user.name}</h3>
              <p>{user.email}</p>
              <p>{user.address.city}</p>
            </div>
          ))}
        </div>
      </div>


    </>
  )

}