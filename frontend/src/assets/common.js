export default function errorHandler (err) {
  console.log(err);
  let errStatusCode = err.response.status
  if (errStatusCode === 404) {
      alert(err.message + "데이터를 불러오는 도중 오류가 발생하였습니다 관리자에게 문의하십시오")
      return
  }

  if (errStatusCode === 403 || errStatusCode === 401) {
      this.$router.push("/login")
      return
  }
  alert("통신중 오류가 발생하였습니다 error: ", err)
}