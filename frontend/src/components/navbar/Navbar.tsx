import React from "react";
import Link from "next/link";

const Navbar = () => {
  return (
    <nav className="flex px-20 py-10">
      <div>Ecommerce</div>
      <div className="flex-1 flex justify-center">Search</div>

      <ul className="flex gap-4">
        <li>
          <Link href="/user">User</Link>
        </li>
        <li>
          <Link href="/signup">Sign up</Link>
        </li>
        <li>
          <Link href="/delete">Delete Account</Link>
        </li>
        <li>
          <Link href="/delete-permanently">Delete Permanently</Link>
        </li>
      </ul>
    </nav>
  );
};

export default Navbar;
