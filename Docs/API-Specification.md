# SIGNCONNECT API SPECIFICATION

## 1. Signup

- Endpoint POST /auth/signup
- Request Headers :

  1. A-AUTHORIZATION : API_KEY

- Request Expectation :

```json
{
  "username": "jono",
  "email": "jono@testing.com",
  "password": "testingtesting"
}
```

- Respond Expectation :

```json
{
  "id": 1,
  "email": "jono@testing.com"
}
```

## 2. Signin

- Endpoint POST /auth/signin
- Request Headers :

  1. A-AUTHORIZATION : API_KEY

- Request Expectation :

```json
{
  "email": "jono@testing.com",
  "password": "testingtesting"
}
```

- Respond Expectation :

```json
{
    "id" : 1,
    "token" : {UNIQE_TOKEN}
}
```

## 3. Get 10 player Leaderboard

- Endpoint GET /leaderboard?limit=10
- Request Headers :

  1. A-AUTHORIZATION : API_KEY
  2. Cookie (X-AUTHORIZATION) : JWT_TOKEN (Login First)

- Request Expectation :

```json
{}
```

- Respond Expectation :

```json
{
    "leaderboard" : [
        {
            "id" : 5,
            "username" : "jono1",
            "experience" : 2341
        },
        {
            "id" : 1,
            "username" : "jono3",
            "experience" : 2141
        },
        {
            "id" : 3,
            "username" : "jono8",
            "experience" : 2003
        },
        {
            "id" : 2,
            "username" : "jono2",
            "experience" : 1732
        },
        {
            "id" : 9,
            "username" : "jono9",
            "experience" : 1576
        },

        ... (10 User Data)
    ]
}
```

## 4. Get user Profile

- Endpoint GET /user/profile
- Request Headers :

  1. A-AUTHORIZATION : API_KEY
  2. cookie (X-AUTHORIZATION) : JWT_TOKEN (Login First)

- Request Expectation :

```json
{}
```

- Respond Expectation :

```json
{
  "user": {
    "id": 1,
    "username": "jono",
    "email": "jono@testing.com",
    "role": "user",
    "level": 5,
    "experience": 1000
  }
}
```

## 5. Get Questions By Level

- Endpoint GET /game/questions/:level
- Request Headers :

  1. A-AUTHORIZATION : API_KEY
  2. cookie (X-AUTHORIZATION) : JWT_TOKEN (Login First)

- Request Expectation :

```json
{}
```

- Respond Expectation :

```json
{
  "result": [
    {
        "id" : 1,
        "level" : 1,
        "questions_image" : null,
        "questions" : "Gambar yang benar untuk huruf \"A\" pada bahasa isyarat adalah...",
        "ans1" : {"image_url"},
        "ans2" : {"image_url"},
        "ans3" : {"image_url"},
        "ans4" : {"image_url"},
        "correct_ans" : {"image_url"}
    },
    {
        "id" : 1,
        "level" : 1,
        "questions_image" : {"image_url"},
        "questions" : "Gambar diatas merepresentasikan huruf alfabet apa?",
        "ans1" : "S",
        "ans2" : "A",
        "ans3" : "C",
        "ans4" : "J",
        "correct_ans" : "A"
    },
    {
        "id" : 1,
        "level" : 1,
        "questions_image" : null,
        "questions" : "Gambar yang benar untuk huruf \"C\" pada bahasa isyarat adalah...",
        "ans1" : {"image_url"},
        "ans2" : {"image_url"},
        "ans3" : {"image_url"},
        "ans4" : {"image_url"},
        "correct_ans" : {"image_url"}
    },
    ... (10 Questions Respond)
  ]
}
```

## 6. Post Answer Questions By Level

- Endpoint POST /game/answer/:level
- Request Headers :

  1. A-AUTHORIZATION : API_KEY
  2. cookie (X-AUTHORIZATION) : JWT_TOKEN (Login First)

- Request Expectation :

```json
{
  "user_answer" : [
    {
        "id" : 1,
        "user_ans" : {"image_url"},
        "correct_ans" : {"image_url"}
    },
    {
        "id" : 2,
        "user_ans" : "A",
        "correct_ans" : "A"
    },
    {
        "id" : 3,
        "user_ans" : {"image_url"},
        "correct_ans" : {"image_url"}
    },
  ]
}
```

- Respond Expectation :

```json
{
  "result": {
    "level": 1,
    "experience_earned": 85
  }
}
```

## 7. Get Dictionary Kategori

- Endpoint GET /dictionary/kategori
- Request Headers :

  1. A-AUTHORIZATION : API_KEY
  2. cookie (X-AUTHORIZATION) : JWT_TOKEN (Login First)

- Request Expectation :

```json
{}
```

- Respond Expectation :

```json
{
  "result": [
    {
      "id": 1,
      "kategori": "Alfabet"
    },
    {
      "id": 2,
      "kategori": "Kata Dasar"
    },
    {
      "id": 3,
      "kategori": "Perkenalan"
    }
  ]
}
```

## 7. Get Dictionary List By Kategori

- Endpoint GET /dictionary/:kategori/list
- Request Headers :

  1. A-AUTHORIZATION : API_KEY
  2. cookie (X-AUTHORIZATION) : JWT_TOKEN (Login First)

- Request Expectation :

```json
{}
```

- Respond Expectation :

```json
{
  "result": [
    {
      "id": 1,
      "kategori_id": "1",
      "huruf": "A"
    },
    {
      "id": 2,
      "kategori_id": "1",
      "huruf": "B"
    },
    {
      "id": 3,
      "kategori_id": "1",
      "huruf": "C"
    }
  ]
}
```

## 8. Get Dictionary Value By Kategori & id

- Endpoint GET /dictionary/:kategori/:value_id
- Request Headers :

  1. A-AUTHORIZATION : API_KEY
  2. cookie (X-AUTHORIZATION) : JWT_TOKEN (Login First)

- Request Expectation :

```json
{}
```

- Respond Expectation :

```json
{
  "result": [
    {
      "id": 1,
      "kategori_id": "1",
      "huruf": "A",
      "image_url" : {"image_url"}
    },
  ]
}
```

## 9. Get Dictionary Value By Kategori & id

- Endpoint GET /dictionary/:kategori/:value_id
- Request Headers :

  1. A-AUTHORIZATION : API_KEY
  2. cookie (X-AUTHORIZATION) : JWT_TOKEN (Login First)

- Request Expectation :

```json
{}
```

- Respond Expectation :

```json
{
  "result": [
    {
      "id": 1,
      "kategori_id": "1",
      "huruf": "A",
      "image_url" : {"image_url"}
    },
  ]
}
```

## 10. Get Daftar Lembaga Disabilitas

- Endpoint GET /lembaga/
- Request Headers :

  1. A-AUTHORIZATION : API_KEY
  2. cookie (X-AUTHORIZATION) : JWT_TOKEN (Login First)

- Request Expectation :

```json
{}
```

- Respond Expectation :

```json
{
  "result": [
    {
      "id": 1,
      "nama": "Pusat Layanan Juru Bahasa Isyarat",
      "lokasi": "Malang, Jawa Timur",
      "lembaga_image_url" : {"image_url"},
      "minimal_tahun" : 1
    },
    {
      "id": 2,
      "nama": "Pusat Bahasa Isyarat Indonesia",
      "lokasi": "Malang, Jawa Timur",
      "lembaga_image_url" : {"image_url"},
      "minimal_tahun" : 1
    },
    {
      "id": 3,
      "nama": "Akar Tuli Malang",
      "lokasi": "Malang, Jawa Timur",
      "lembaga_image_url" : {"image_url"},
      "minimal_tahun" : 0
    },
  ]
}
```

## 11. Get Detail Lembaga Disabilitas

- Endpoint GET /lembaga/:id
- Request Headers :

  1. A-AUTHORIZATION : API_KEY
  2. cookie (X-AUTHORIZATION) : JWT_TOKEN (Login First)

- Request Expectation :

```json
{}
```

- Respond Expectation :

```json
{
  "result": {
      "id": 1,
      "nama": "Pusat Layanan Juru Bahasa Isyarat",
      "lokasi": "Malang, Jawa Timur",
      "lembaga_image_url" : {"image_url"},
      "minimal_tahun" : 1,
      "tentang_lembaga" : "text....",
      "deskripsi_pekerjaan" : "text....",
      "persyaratan" : "text...."
    }
}
```

## 12. Send Lamaran Volunteer ke Lembaga

- Endpoint POST /lembaga/:id
- Request Headers :

  1. A-AUTHORIZATION : API_KEY
  2. cookie (X-AUTHORIZATION) : JWT_TOKEN (Login First)

- Request Expectation :

```json
{
    "data" : {
        "cv_url" : {"url"},
        "username" : {"user.username"},
        "email" : {"user.email"},
        "lembaga_id" : 1
    }
}
```

- Respond Expectation :

```json
{
  "result": {
    "lembaga_id": 1,
    "massage": "Success Mengirimkan Lamaran"
  }
}
```
