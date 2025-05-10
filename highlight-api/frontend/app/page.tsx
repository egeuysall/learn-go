export default function Home() {
  return (
      <main className="grid place-items-center h-screen w-screen">
          <section className="grid gap-4 w-[90vw]" >
            <h1 className="text-2xl md:text-3xl lg:text-3xl font-bold text-center">Welcome to the API!</h1>
              <p className="text-center">Please go to <code>/add</code> to add elements and <code>/see</code> to see elements.</p>
          </section>
      </main>
  );
}
