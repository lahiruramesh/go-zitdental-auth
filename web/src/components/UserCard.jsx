export function UserCard({ user, onSelect }) {
    return (
      <div 
        onClick={() => onSelect(user)}
        className="rounded-xl cursor-pointer transform transition-all duration-300 hover:scale-105 border-gray shadow-md hover:shadow-xl"
      >
        <div className="bg-white p-6">
          <div className="border-b pb-4 mb-4">
            <h3 className="text-xl font-semibold text-gray-800">
              {user.firstName} {user.lastName}
            </h3>
            <p className="text-gray-500 text-sm mt-1">{user.roles[0]}</p>
          </div>
        </div>
      </div>
    );
  }