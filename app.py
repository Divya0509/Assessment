from flask import Flask, jsonify, render_template, request, redirect, url_for, flash
import psycopg2
from psycopg2 import IntegrityError, DataError, Error

app = Flask(__name__)
app.secret_key = 'my_key'

def get_db_connection():
    conn = psycopg2.connect(
        host="localhost",
        database="movie_db",
        user="postgres",
        password="System@1234"
    )
    return conn

@app.route('/movies', methods=['GET'])
def get_movies():
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute('SELECT * FROM movies')
    movies = cursor.fetchall()
    cursor.close()
    conn.close()
    
    movies_list = []
    for movie in movies:
        movies_list.append({
            'id': movie[0],
            'title': movie[1],
            'description': movie[2],
            'year': movie[3],
            'rating': movie[4],
            'genre': movie[5]
        })
    
    return jsonify(movies_list)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/add')
def add_movie_page():
    return render_template('add_movie.html')

@app.route('/add_movie', methods=['POST'])
def add_movie():
    title = request.form['title']
    description = request.form['description']
    year = request.form['year']
    rating = request.form['rating']
    genre = request.form['genre']
    
    conn = get_db_connection()
    cursor = conn.cursor()
    try:
        cursor.execute(
            'INSERT INTO movies (title, description, year, ratings, genre) VALUES (%s, %s, %s, %s, %s)',
            (title, description, year, rating, genre)
        )
        conn.commit()
        flash('Movie added successfully!!!', 'success')
    except (IntegrityError, DataError, Error) as e:
        conn.rollback()
        flash(f'Error adding movie: {e.pgerror}')
        flash(f'Error adding movie: {e.pgerror}', 'danger')

    finally:
        cursor.close()
        conn.close()
    
    return redirect(url_for('index'))

@app.route('/edit_movie/<int:id>', methods=['GET', 'POST'])
def edit_movie(id):
    conn = get_db_connection()
    cursor = conn.cursor()
    
    if request.method == 'POST':
        description = request.form['description']
        year = request.form['year']
        ratings = request.form['ratings']
        genre = request.form['genre']
        
        cursor.execute('UPDATE movies SET description = %s, year = %s, ratings = %s, genre = %s WHERE id = %s',(description, year, ratings, genre, id))
        conn.commit()
        cursor.close()
        conn.close()
        flash('Movie updated successfully!', 'success')
        return redirect(url_for('index'))
    
    cursor.execute('SELECT * FROM movies WHERE id = %s', (id,))
    movie = cursor.fetchone()
    cursor.close()
    conn.close()
    
    return render_template('edit_movie.html', movie=movie)


# @app.route('/delete_movie/<int>', methods=['POST'])
@app.route('/delete_movie/<int:id>', methods=['POST'])
def delete_movie(id):
    conn = get_db_connection()
    cursor = conn.cursor()
    try:
        cursor.execute('DELETE FROM movies WHERE id = %s', (id,))
        conn.commit()
        flash('Movie deleted successfully!!!', 'success')
    except Error as e:
        conn.rollback()
        flash(f'Error deleting movie: {e.pgerror}', 'danger')
    finally:
        cursor.close()
        conn.close()
    
    return redirect(url_for('index'))

if __name__ == '__main__':
    app.run(debug=True)
